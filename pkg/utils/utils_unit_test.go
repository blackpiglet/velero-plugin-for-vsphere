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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"

	"github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/constants"
	veleroplugintest "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/test"
)

func TestGetStringFromParamsMap(t *testing.T) {
	params := make(map[string]interface{})
	params["ValidKey"] = "ValidValue"
	params["NonString"] = false
	tests := []struct {
		name          string
		key           string
		expectedValue string
		ok            bool
	}{
		{
			name:          "Valid key with string value should return corresponding value and true",
			key:           "ValidKey",
			expectedValue: "ValidValue",
			ok:            true,
		},
		{
			name:          "If value in map is non-string should return empty string and false",
			key:           "NonString",
			expectedValue: "",
			ok:            false,
		},
		{
			name:          "No such key should return empty string and false",
			key:           "NoSuchKey",
			expectedValue: "",
			ok:            false,
		},
	}

	logger := veleroplugintest.NewLogger()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str, ok := GetStringFromParamsMap(params, test.key, logger)
			assert.Equal(t, test.expectedValue, str)
			assert.Equal(t, test.ok, ok)
		})
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name        string
		str         string
		defValue    bool
		expectedVal bool
	}{
		{
			name:        "Pos1",
			str:         "true",
			defValue:    false,
			expectedVal: true,
		},
		{
			name:        "Pos2",
			str:         "false",
			defValue:    true,
			expectedVal: false,
		},
		{
			name:        "Empty string",
			str:         "",
			defValue:    false,
			expectedVal: false,
		},
		{
			name:        "Invalid str",
			str:         "AAA",
			defValue:    true,
			expectedVal: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := GetBool(test.str, test.defValue)
			require.Equal(t, test.expectedVal, res)
		})
	}
}

func TestGetComponentFromImage(t *testing.T) {
	tests := []struct {
		name                                             string
		image                                            string
		expectedRepo, expectedContainer, expectedVersion string
	}{
		{
			name:              "ExpectedDummyCase",
			image:             "a/b/c/d:x",
			expectedRepo:      "a/b/c",
			expectedContainer: "d",
			expectedVersion:   "x",
		},
		{
			name:              "ExpectedDockerhubCase",
			image:             "vsphereveleroplugin/velero-plugin-for-vsphere:1.1.0-rc2",
			expectedRepo:      "vsphereveleroplugin",
			expectedContainer: "velero-plugin-for-vsphere",
			expectedVersion:   "1.1.0-rc2",
		},
		{
			name:              "ExpectedCustomizedCase",
			image:             "xyz-repo.vmware.com/velero/velero-plugin-for-vsphere:1.1.0-rc2",
			expectedRepo:      "xyz-repo.vmware.com/velero",
			expectedContainer: "velero-plugin-for-vsphere",
			expectedVersion:   "1.1.0-rc2",
		},
		{
			name:              "ExpectedLocalCase",
			image:             "velero-plugin-for-vsphere:1.1.0-rc2",
			expectedRepo:      "",
			expectedContainer: "velero-plugin-for-vsphere",
			expectedVersion:   "1.1.0-rc2",
		},
		{
			name:              "ExpectedNonTaggedImageCase",
			image:             "xyz-repo.vmware.com/velero/velero-plugin-for-vsphere",
			expectedRepo:      "xyz-repo.vmware.com/velero",
			expectedContainer: "velero-plugin-for-vsphere",
			expectedVersion:   "",
		},
		{
			name:              "ExpectedCaseOfRegistryEndpointWithPort",
			image:             "xyz-repo.vmware.com:9999/velero/velero-plugin-for-vsphere:1.1.0-rc2",
			expectedRepo:      "xyz-repo.vmware.com:9999/velero",
			expectedContainer: "velero-plugin-for-vsphere",
			expectedVersion:   "1.1.0-rc2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualRepo := GetComponentFromImage(test.image, constants.ImageRepositoryComponent)
			actualContainer := GetComponentFromImage(test.image, constants.ImageContainerComponent)
			actualVersion := GetComponentFromImage(test.image, constants.ImageVersionComponent)
			assert.Equal(t, actualRepo, test.expectedRepo)
			assert.Equal(t, actualContainer, test.expectedContainer)
			assert.Equal(t, actualVersion, test.expectedVersion)
		})
	}
}

func TestGetVersionFromImage(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		containers []corev1.Container
		expected   string
	}{
		{
			name: "Valid image string should return non-empty version",
			key:  "cloud-provider-vsphere/csi/release/driver",
			containers: []corev1.Container{
				{
					Image: "gcr.io/cloud-provider-vsphere/csi/release/driver:corev1.0.1",
				},
			},
			expected: "corev1.0.1",
		},
		{
			name: "Valid image string should return non-empty version",
			key:  "cloud-provider-vsphere/csi/release/driver",
			containers: []corev1.Container{
				{
					Image: "cloud-provider-vsphere/csi/release/driver:v2.0.0",
				},
			},
			expected: "v2.0.0",
		},
		{
			name: "Valid image string should return non-empty version",
			key:  "cloud-provider-vsphere/csi/release/driver",
			containers: []corev1.Container{
				{
					Image: "myregistry/cloud-provider-vsphere/csi/release/driver:v2.0.0",
				},
			},
			expected: "v2.0.0",
		},
		{
			name: "Valid image string should return non-empty version",
			key:  "cloud-provider-vsphere/csi/release/driver",
			containers: []corev1.Container{
				{
					Image: "myregistry/level1/level2/cloud-provider-vsphere/csi/release/driver:v2.0.0",
				},
			},
			expected: "v2.0.0",
		},
		{
			name: "Invalid image name should return empty string",
			key:  "cloud-provider-vsphere/csi/release/driver",
			containers: []corev1.Container{
				{
					Image: "gcr.io/csi/release/driver:corev1.0.1",
				},
			},
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			version := GetVersionFromImage(test.containers, test.key)
			assert.Equal(t, test.expected, version)
		})
	}
}
