package util

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/constants"
	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/utils"
)

// RemoveAnnotations removes the supplied keys from the annotations on the object
func RemoveAnnotations(o *metav1.ObjectMeta, keys []string) {
	if o.Annotations != nil {
		for _, k := range keys {
			delete(o.Annotations, k)
		}
	}
}

// AddAnnotations adds the supplied key-values to the annotations on the object
func AddAnnotations(o *metav1.ObjectMeta, vals map[string]string) {
	if o.Annotations == nil {
		o.Annotations = make(map[string]string)
	}
	for k, v := range vals {
		o.Annotations[k] = v
	}
}

/*
 * Extracts the CRD name from a K8S Unstructured into the CRD name.  Supports namespaced and cluster level CRs
 * K8S Cluster Resource Self Link format: /api/<version>/<resource plural name>/<item name>,
 *     e.g. /api/v1/persistentvolumes/pvc-3240e5ed-9a97-446c-a6ab-b2442d852d04
 * K8S Resource Name = <resource plural name>, e.g. persistentvolumes
 * Custom Resource Cluster Self Link format: /apis/<CR group>/<version>/<CR plural name>/<item name>,
 *     e.g. /api/cnsdp.vmware.com/v1/backuprepositories/br-1
 * Custom Resource Name = <CR plural name>.<CR group>, e.g. backuprepositories.cnsdp.vmware.com
 * TODO - replace with a better system that does need singular to plural translation
 */
func UnstructuredToCRDName(item runtime.Unstructured) (string, error) {
	// Retrieve the common object metadata and check to see if this is a blocked type
	accessor := meta.NewAccessor()
	apiVersion, err := accessor.APIVersion(item)
	if err != nil {
		return "", err
	}
	var group string

	gv, err := schema.ParseGroupVersion(apiVersion)
	kind, err := accessor.Kind(item)
	if err != nil {
		return "", err
	}
	group = gv.Group

	var pluralName string
	// Singular to plural conversion is done by adding an 's' unless the kind ends in a 'y' in which case
	// y gets replaced with ies.  This is not a complete set of rules, it doesn't handle shelf->shelves for example
	// but the rules are weird.  Currently works for the resources defined, this is the place to add additional plural rules
	// if necessary
	if strings.HasSuffix(kind, "y") {
		pluralName = strings.ToLower(kind)[0:len(kind)-1] + "ies"
	} else {
		pluralName = strings.ToLower(kind) + "s"
	}
	if group != "" {
		return pluralName + "." + group, nil
	}
	return pluralName, nil
}

func GetKubeClient(config *rest.Config, logger logrus.FieldLogger) (*kubernetes.Clientset, error) {
	var err error
	if config == nil {
		config, err = utils.GetKubeClientConfig()
		if err != nil {
			logger.WithError(err).Errorf("Failed to get k8s inClusterConfig")
			return nil, errors.Wrap(err, "could not retrieve in-cluster config")
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.WithError(err).Error("Failed to get k8s clientset from the given config")
		return nil, err
	}
	return clientset, nil
}

func IsObjectBlocked(item runtime.Unstructured, logger logrus.FieldLogger) (bool, string, error) {
	crdName, err := UnstructuredToCRDName(item)
	if err != nil {
		return false, "", errors.Errorf("Could not translate item kind %s to CRD name", item.GetObjectKind())
	}

	blockListConfigMap, err := RetrieveBlockListConfigMap(logger)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return false, crdName, errors.WithStack(err)
		}
		return false, crdName, nil
	}
	if IsResourceBlocked(crdName, blockListConfigMap) {
		return true, crdName, nil
	}
	return false, crdName, nil
}

func GetResources(blockListConfigMap map[string]string) []string {
	desiredResources := make([]string, len(constants.ResourcesToHandle)+len(blockListConfigMap))
	for resourceToHandle, _ := range constants.ResourcesToHandle {
		desiredResources = append(desiredResources, resourceToHandle)
	}
	for resourceToBlock, _ := range blockListConfigMap {
		desiredResources = append(desiredResources, resourceToBlock)
	}
	return desiredResources
}

func IsResourceBlocked(resourceName string, resourcesToBlock map[string]string) bool {
	isBlocked, _ := strconv.ParseBool(resourcesToBlock[resourceName])
	return isBlocked
}

func IsResourceBlockedOnRestore(resourceName string) bool {
	return constants.ResourcesToBlockOnRestore[resourceName]
}

func Contains(slice []string, key string) bool {
	for _, i := range slice {
		if i == key {
			return true
		}
	}
	return false
}

// If there is no configmap velero-vsphere-plugin-block-list created before, create a new configmap from the default blocking list.
// If there is already a configmap velero-vsphere-plugin-block-list created, leave the previous one unchanged.

func CreateBlockListConfigMap(logger logrus.FieldLogger) error {
	veleroNs, exist := os.LookupEnv("VELERO_NAMESPACE")
	if !exist {
		errMsg := "Failed to lookup the ENV variable for velero namespace"
		return errors.New(errMsg)
	}

	restConfig, err := utils.GetKubeClientConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return errors.WithStack(err)
	}

	ctx := context.Background()
	var create bool
	blockListConfigMap, err := kubeClient.CoreV1().ConfigMaps(veleroNs).Get(ctx, constants.ResourcesToBlockListName, metav1.GetOptions{})
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			logger.Errorf("Failed to retrieve %s configuration.\n", constants.ResourcesToBlockListName)
			return err
		}
		blockListConfigMap = &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      constants.ResourcesToBlockListName,
				Namespace: veleroNs,
			},
			Data: make(map[string]string),
		}
		create = true
	}

	// If there is blocking list configmap created already, leave the previous one unchanged.
	if !create {
		logger.Info("Config map that blocks resources for backup/restore already exists. Using the existing one.")
		return nil
	}

	blockListData := make(map[string]string)
	for resourceToBlock, isBlock := range constants.ResourcesToBlock {
		blockListData[resourceToBlock] = strconv.FormatBool(isBlock)
	}
	blockListConfigMap.Data = blockListData
	_, err = kubeClient.CoreV1().ConfigMaps(veleroNs).Create(ctx, blockListConfigMap, metav1.CreateOptions{})
	if err != nil {
		logger.Errorf("failed to create the config map %s that blocks resources for backup and restore.\n", constants.ResourcesToBlockListName)
		return err
	}

	return nil
}

func RetrieveBlockListConfigMap(logger logrus.FieldLogger) (map[string]string, error) {
	ctx := context.Background()

	veleroNs, exist := os.LookupEnv("VELERO_NAMESPACE")
	if !exist {
		errMsg := "Failed to lookup the ENV variable for velero namespace"
		return nil, errors.New(errMsg)
	}

	restConfig, err := utils.GetKubeClientConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	blockListConfigMap, err := kubeClient.CoreV1().ConfigMaps(veleroNs).Get(ctx, constants.ResourcesToBlockListName, metav1.GetOptions{})
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			logger.Errorf("Failed to retrieve ConfigMap: %s. Error: %v\n", constants.ResourcesToBlockListName, err)
		} else {
			logger.Errorf("Failed to find ConfigMap: %s. Error: %v\n", constants.ResourcesToBlockListName, err)
		}
		return nil, err
	}

	return blockListConfigMap.Data, nil
}
