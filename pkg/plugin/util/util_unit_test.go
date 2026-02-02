package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_ItemToCRDName(t *testing.T) {
	accessor := meta.NewAccessor()

	backupUnstructuredMock := unstructured.Unstructured{}
	accessor.SetKind(&backupUnstructuredMock, "Backup")
	accessor.SetAPIVersion(&backupUnstructuredMock, "velero.io/v1")

	backupCRDName, err := UnstructuredToCRDName(&backupUnstructuredMock)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, backupCRDName, "backups.velero.io")

	repositoryUnstructuredMock := unstructured.Unstructured{}
	accessor.SetKind(&repositoryUnstructuredMock, "ResticRepository")
	accessor.SetAPIVersion(&repositoryUnstructuredMock, "velero.io/v1")

	repositoryCRDName, err := UnstructuredToCRDName(&repositoryUnstructuredMock)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, repositoryCRDName, "resticrepositories.velero.io")

	pvcUnstructuredMock := unstructured.Unstructured{}
	accessor.SetKind(&pvcUnstructuredMock, "PersistentVolumeClaim")
	accessor.SetAPIVersion(&pvcUnstructuredMock, "v1")

	pvcCRDName, err := UnstructuredToCRDName(&pvcUnstructuredMock)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, pvcCRDName, "persistentvolumeclaims")
}

var (
	csiStorageClass = "vsphere-csi-sc"
)

func TestContains(t *testing.T) {
	testCases := []struct {
		name           string
		inSlice        []string
		inKey          string
		expectedResult bool
	}{
		{
			name:           "should find the key",
			inSlice:        []string{"key1", "key2", "key3", "key4", "key5"},
			inKey:          "key3",
			expectedResult: true,
		},
		{
			name:           "should not find the key in non-empty slice",
			inSlice:        []string{"key1", "key2", "key3", "key4", "key5"},
			inKey:          "key300",
			expectedResult: false,
		},
		{
			name:           "should not find key in empty slice",
			inSlice:        []string{},
			inKey:          "key300",
			expectedResult: false,
		},
		{
			name:           "should not find key in nil slice",
			inSlice:        nil,
			inKey:          "key300",
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := Contains(tc.inSlice, tc.inKey)
			assert.Equal(t, tc.expectedResult, actualResult)
		})
	}
}

func TestIsResourceBlocked(t *testing.T) {
	var resourcesToBlockForTest = map[string]string{
		"agentinstalls.installers.tmc.cloud.vmware.com": "true",
		"availabilityzones.topology.tanzu.vmware.com":   "false",
	}

	testCases := []struct {
		name            string
		crdName         string
		resourceToBlock map[string]string
		expectedResult  bool
	}{
		{
			name:            "crd is blocked",
			crdName:         "agentinstalls.installers.tmc.cloud.vmware.com",
			resourceToBlock: resourcesToBlockForTest,
			expectedResult:  true,
		},
		{
			name:            "crd is not blocked",
			crdName:         "availabilityzones.topology.tanzu.vmware.com",
			resourceToBlock: resourcesToBlockForTest,
			expectedResult:  false,
		},
		{
			name:            "crd is not included in skip list",
			crdName:         "donotexist",
			resourceToBlock: resourcesToBlockForTest,
			expectedResult:  false,
		},
		{
			name:            "block list is nil",
			crdName:         "agentinstalls.installers.tmc.cloud.vmware.com",
			resourceToBlock: nil,
			expectedResult:  false,
		},
		{
			name:            "block list is empty",
			crdName:         "availabilityzones.topology.tanzu.vmware.com",
			resourceToBlock: make(map[string]string),
			expectedResult:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isBlocked := IsResourceBlocked(tc.crdName, tc.resourceToBlock)
			assert.Equal(t, tc.expectedResult, isBlocked)
		})
	}
}
