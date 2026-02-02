/*
 * Copyright 2019 the Velero contributors
 * SPDX-License-Identifier: Apache-2.0
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	// enable fips only mode
	_ "crypto/tls/fipsonly"

	"github.com/sirupsen/logrus"
	velero_plugin "github.com/vmware-tanzu/velero/pkg/plugin/framework"

	plugins_pkg "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/plugin"
	plugins_util "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/plugin/util"
)

func main() {
	logger := logrus.StandardLogger()

	if err := plugins_util.CreateBlockListConfigMap(logger); err != nil {
		logrus.Fatalf("Failed to create block list config map: %v", err)
	}

	veleroPluginServer := velero_plugin.NewServer()
	veleroPluginServer = veleroPluginServer.
		RegisterBackupItemAction("velero.io/vsphere-backupper", newBackupItemAction).
		RegisterRestoreItemAction("velero.io/vsphere-restorer", newRestoreItemAction)
	veleroPluginServer.Serve()
}

func newBackupItemAction(logger logrus.FieldLogger) (interface{}, error) {
	return &plugins_pkg.NewBackupItemAction{Log: logger}, nil
}

func newRestoreItemAction(logger logrus.FieldLogger) (interface{}, error) {
	return &plugins_pkg.NewRestoreItemAction{Log: logger}, nil
}
