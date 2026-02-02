/*
Copyright 2020 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	k8sv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/constants"
)

func GetStringFromParamsMap(params map[string]interface{}, key string, logger logrus.FieldLogger) (value string, ok bool) {
	valueIF, ok := params[key]
	if ok {
		value, ok := valueIF.(string)
		if !ok {
			logger.Errorf("Value for params key %s is not a string", key)
		}
		return value, ok
	} else {
		logger.Infof("No such key %s in params map", key)
		return "", ok
	}
}

func GetBool(str string, defValue bool) bool {
	if str == "" {
		return defValue
	}

	res, err := strconv.ParseBool(str)
	if err != nil {
		res = defValue
		err = nil
	}

	return res
}

type NotFoundError struct {
	errMsg string
}

func (this NotFoundError) Error() string {
	return this.errMsg
}

func NewNotFoundError(errMsg string) NotFoundError {
	err := NotFoundError{
		errMsg: errMsg,
	}
	return err
}

func GetComponentFromImage(image string, component string) string {
	components := GetComponentsFromImage(image)
	return components[component]
}

func GetComponentsFromImage(image string) map[string]string {
	components := make(map[string]string)

	if image == "" {
		return components
	}

	var taggedContainer string
	lastIndex := strings.LastIndex(image, "/")
	if lastIndex < 0 {
		taggedContainer = image
	} else {
		components[constants.ImageRepositoryComponent] = image[:lastIndex]
		taggedContainer = image[lastIndex+1:]
	}

	parts := strings.SplitN(taggedContainer, ":", 2)
	if len(parts) == 2 {
		components[constants.ImageVersionComponent] = parts[1]
	}
	components[constants.ImageContainerComponent] = parts[0]

	return components
}

type ClientConfigNotFoundError struct {
	errMsg string
}

func (this ClientConfigNotFoundError) Error() string {
	return this.errMsg
}

func NewClientConfigNotFoundError(errMsg string) ClientConfigNotFoundError {
	err := ClientConfigNotFoundError{
		errMsg: errMsg,
	}
	return err
}

func GetKubeClientConfig() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here

	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	clientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Error finding Kubernetes API server config in $KUBECONFIG, or in-cluster configuration")
	}

	return clientConfig, nil
}

func CreateKubeClientSet() (*kubernetes.Clientset, error) {
	clientConfig, err := GetKubeClientConfig()
	if err != nil {
		return nil, NewClientConfigNotFoundError(fmt.Sprintf("Could not get client config with err: %v", err))
	}

	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Could not create kubernetes clientset")
	}
	return clientset, err
}

func GetCSIInstalledVersion(kubeClient kubernetes.Interface) (string, error) {
	csiVersion := ""
	// For CSI >=2.3.0, the CSI Deployment is in vmware-system-csi namespace.
	csiDeployment, err := kubeClient.AppsV1().Deployments(constants.VSphereCSIControllerNamespace).Get(context.TODO(), constants.VSphereCSIController, metav1.GetOptions{})
	if err == nil {
		if csiVersion, err = GetCSIVersionFromImage(csiDeployment.Spec.Template.Spec.Containers); err != nil {
			// Image not found.
			return "", err
		}
		if csiVersion == constants.CsiDevVersion {
			// Using a unreleased csi version, defaulting to 2.3.0 since vSphere csi driver 2.3.0
			// is always installed in vmware-system-csi namespace.
			return constants.Csi2_3_0_Version, nil
		}
		return csiVersion, nil
	}

	// Search in kube-system next.
	csiDeployment, err = kubeClient.AppsV1().Deployments(constants.KubeSystemNamespace).Get(context.TODO(), constants.VSphereCSIController, metav1.GetOptions{})
	if err == nil {
		if csiVersion, err = GetCSIVersionFromImage(csiDeployment.Spec.Template.Spec.Containers); err != nil {
			// Image not found.
			return "", err
		}
		if csiVersion == constants.CsiDevVersion {
			// Using unreleased version, defaulting to 2.1.0 since vSphere CSI driver is installed as a
			// Deployment on `kube-system` for versions greater than v1.0.3 and lesser than v2.3.0
			return constants.Csi2_0_0_Version, nil
		}
		return csiVersion, nil
	}
	// CSI driver is deployed as StatefulSet for v1.0.2 and v1.0.3.
	csiStatefulset, err := kubeClient.AppsV1().StatefulSets(constants.KubeSystemNamespace).Get(context.TODO(), constants.VSphereCSIController, metav1.GetOptions{})
	if err == nil {
		if csiVersion, err = GetCSIVersionFromImage(csiStatefulset.Spec.Template.Spec.Containers); err != nil {
			// Image not found.
			return "", err
		}
		if csiVersion == constants.CsiDevVersion {
			// Using unreleased version, defaulting to 1.0.2 since vSphere CSI driver is installed as a
			// StatefulSet on `kube-system` namespace for versions greater than v1.0.2 and lesser than v2.0.0
			return constants.CsiMinVersion, nil
		}
		return csiVersion, nil
	}

	// vSphere CSI controller not installed.
	return csiVersion, errors.Errorf("vSphere CSI controller, %s, is required by velero-plugin-for-vsphere. Please make sure the vSphere CSI controller is installed in the cluster", constants.VSphereCSIController)
}

func GetCSIVersionFromImage(containers []k8sv1.Container) (string, error) {
	var csiDriverVersion string
	csiDriverVersion = GetVersionFromImage(containers, "cloud-provider-vsphere/csi/release/driver")
	if csiDriverVersion == "" {
		csiDriverVersion = GetVersionFromImage(containers, "cloud-provider-vsphere/csi/ci/driver")
		// The version received from cloud-provider-vsphere/csi/ci/driver is not nil.
		// This is used to support latest images on public repository.
		if csiDriverVersion != "" {
			// Developer image
			csiDriverVersion = constants.CsiDevVersion
		}
	}
	if csiDriverVersion == "" {
		// vSphere driver found but the images are invalid.
		return "", errors.New("Expected CSI driver images not found")
	}
	return csiDriverVersion, nil
}

func GetCSIClusterTypeFromEnv(containers []k8sv1.Container) (constants.ClusterFlavor, error) {
	for _, container := range containers {
		if container.Name == constants.VSphereCSIController {
			// Iterate through the env variables and check if "CLUSTER_FLAVOR" env variable is defined.
			for _, envVar := range container.Env {
				if envVar.Name == "CLUSTER_FLAVOR" {
					if envVar.Value == "WORKLOAD" {
						return constants.Supervisor, nil
					} else if envVar.Value == "GUEST_CLUSTER" {
						return constants.TkgGuest, nil
					} else {
						// Should not happen, if "CLUSTER_FLAVOR" is defined, the value should be guest or supervisor.
						return constants.Unknown, nil
					}
				}
			}
			// For Vanilla deployment the "CLUSTER_FLAVOR" environment variable is not defined.
			return constants.VSphere, nil
		}
	}
	return constants.Unknown, errors.New("Expected CSI driver container images not found while inferring cluster type.")
}

// Return version in the format: vX.Y.Z
func GetVersionFromImage(containers []k8sv1.Container, imageName string) string {
	var tag = ""
	for _, container := range containers {
		if strings.Contains(container.Image, imageName) {
			tag = GetComponentFromImage(container.Image, constants.ImageVersionComponent)
			break
		}
	}
	if tag == "" {
		return ""
	}
	if strings.Contains(tag, "-") {
		imgVersion := strings.Split(tag, "-")[0]
		return imgVersion
	} else {
		return tag
	}
}

func GetKubeClientSet(config *rest.Config) (kubernetes.Interface, error) {
	return kubernetes.NewForConfig(config)
}

func GetVeleroNamespace() (string, bool) {
	return os.LookupEnv("VELERO_NAMESPACE")
}
