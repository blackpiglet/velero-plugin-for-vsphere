package plugin

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/constants"
	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/plugin/util"
	pluginItem "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/plugin/util"
)

// BackupItemAction is a backup item action plugin for Velero.
type NewRestoreItemAction struct {
	Log logrus.FieldLogger
}

// AppliesTo returns information indicating that the BackupItemAction should be invoked to backup resources.
func (p *NewRestoreItemAction) AppliesTo() (velero.ResourceSelector, error) {
	p.Log.Info("VSphere RestoreItemAction AppliesTo")

	blockListConfigMap, err := util.RetrieveBlockListConfigMap(p.Log)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			p.Log.Info("Failed to retrieve resources to block list.")
			return velero.ResourceSelector{}, errors.WithStack(err)
		}
		p.Log.Info("There is no blocklist configmap found. Assuming no resources to block.")
		blockListConfigMap = make(map[string]string)
	}

	resources := []string{"persistentvolumeclaims"}
	for resourceToBlock := range blockListConfigMap {
		resources = append(resources, resourceToBlock)
	}
	for resourceToBlockOnRestore := range constants.ResourcesToBlockOnRestore {
		resources = append(resources, resourceToBlockOnRestore)
	}
	return velero.ResourceSelector{
		IncludedResources: resources,
	}, nil
}

func (p *NewRestoreItemAction) Execute(input *velero.RestoreItemActionExecuteInput) (*velero.RestoreItemActionExecuteOutput, error) {
	blocked, crdName, err := pluginItem.IsObjectBlocked(input.ItemFromBackup, p.Log) // Use ItemFromBackup here so that selflink is available
	if err != nil {
		return nil, errors.Wrap(err, "Failed during IsObjectBlocked check")
	}

	if blocked == false {
		// "pods", "images" and "nsxlbmonitors" are additional resources
		// blocked on restore only for now
		blocked = pluginItem.IsResourceBlockedOnRestore(crdName)
	}
	item := input.Item // Use Item for everything else so that previous actions had a chance to modify the object
	// (e.g. Velero removes extraneous metadata earlier in the restore process)

	p.Log.Infof("Restoring resource %v: blocked = %v", crdName, blocked)

	if blocked {
		if crdName == "pods" {
			return p.createPod(item)
		} else if pluginItem.IsResourceBlockedOnRestore(crdName) {
			// Skip the restore of image and nsxlbmonitor resources on Supervisor Cluster
			p.Log.Infof("Skipping resource %s on restore", crdName)
			return &velero.RestoreItemActionExecuteOutput{
				SkipRestore: true,
			}, nil
		}
		return nil, errors.Errorf("Resource CRD %s is blocked in restore, skipping", crdName)
	}

	var pvc corev1.PersistentVolumeClaim
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(item.UnstructuredContent(), &pvc); err != nil {
		return nil, errors.WithStack(err)
	}

	// Remove the volume health annotations before the restore
	// because vSphere CSI driver has a webhook that prevents this annotation
	// from being created/updated by a non-CSI driver system service account
	healthAnnotations := []string{constants.AnnVolumeHealth, constants.AnnVolumeHealthTS}
	pluginItem.RemoveAnnotations(&pvc.ObjectMeta, healthAnnotations)

	// Update item with the PVC without the volume health annotations
	// so that Velero can create a PVC without being rejected by the
	// vSphere CSI driver webhook in the case of Skipping PVCRestoreItemAction
	pvcMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&pvc)
	if err != nil {
		p.Log.Errorf("Error converting the pvc %s/%s to unstructured. Error: %v", pvc.Namespace, pvc.Name, err)
		return nil, errors.WithStack(err)
	}
	item = &unstructured.Unstructured{Object: pvcMap}

	// exit early if the RestorePVs option is disabled
	p.Log.Info("Checking if the RestorePVs option is disabled in the Restore Spec")
	if input.Restore.Spec.RestorePVs != nil && *input.Restore.Spec.RestorePVs == false {
		p.Log.Infof("Skipping RestoreItemAction for PVC %s/%s since the RestorePVs option is disabled in the Restore Spec.", pvc.Namespace, pvc.Name)
		return &velero.RestoreItemActionExecuteOutput{
			UpdatedItem: item,
		}, nil
	}

	return &velero.RestoreItemActionExecuteOutput{
		SkipRestore: true,
	}, nil
}

func (p *NewRestoreItemAction) createPod(item runtime.Unstructured) (*velero.RestoreItemActionExecuteOutput, error) {
	pod := new(corev1.Pod)
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(item.UnstructuredContent(), &pod); err != nil {
		return nil, errors.WithStack(err)
	}

	if metav1.HasAnnotation(pod.ObjectMeta, constants.VMwareSystemVMUUID) {
		objectMeta := metav1.ObjectMeta{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Labels:    pod.Labels,
		}

		// Only restore Pod annotations not on the list to skip
		if pod.Annotations != nil {
			for annKey, annVal := range pod.Annotations {
				skip := false
				for keyToSkip, _ := range constants.PodAnnotationsToSkip {
					if annKey == keyToSkip {
						// Skip the matching key
						p.Log.Infof("createPod: Skipping annotation %s/%s on Pod %s/%s.", annKey, annVal, pod.Namespace, pod.Name)
						skip = true
						break
					}
				}
				if skip == false {
					metav1.SetMetaDataAnnotation(&objectMeta, annKey, annVal)
					p.Log.Infof("createPod: Set annotation %s:%s on Pod %s/%s.", annKey, annVal, pod.Namespace, pod.Name)
				}
			}

			pod.ObjectMeta = objectMeta
		}

		podMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&pod)
		if err != nil {
			p.Log.Errorf("Error converting the new pod %s/%s to unstructured. Error: %v", pod.Namespace, pod.Name, err)
			return nil, errors.WithStack(err)
		}

		return &velero.RestoreItemActionExecuteOutput{
			UpdatedItem: &unstructured.Unstructured{Object: podMap},
		}, nil
	}

	return &velero.RestoreItemActionExecuteOutput{}, nil
}
