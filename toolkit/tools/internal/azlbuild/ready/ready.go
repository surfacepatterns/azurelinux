// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package ready

import (
	"fmt"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/azlbuildutils"
)

var (
	// get relevant configs
	toolkitDir string
	scriptsDir string
)

// ReadyChanges runs various tools to ready changes for contributing to upstream open source repo
// TODO: use a command builder
func ReadyChanges() (err error) {
//	fmt.Println("[DEBUG] Ready changes")
	azlbuildutils.SetupConfig()
	scriptsDir, _ = azlbuildutils.GetConfig("SCRIPTS_DIR")
	toolkitDir, _ = azlbuildutils.GetConfig("toolkit_root")
//	fmt.Println("[DEBUG] scripts_dir is ", scriptsDir)

	err = checkManifests()
	if err != nil {
		return fmt.Errorf("failed to check manifests:\n%w", err)
	}

/*	err = updateLicenses()
	if err != nil {
		return fmt.Errorf("failed to update licenses:\n%w", err)
	}
*/
	return
}

func specLint() (err error) {
	return
}

// checkManifest runs check_manifests script to spot updates required in manifest files
func checkManifests() (err error) {
	fmt.Println("Checking if manifest scripts are in sync with toolchain specs.\nIf you didn't make changes to any of these specs, the error shown (if any) don't apply to you")
	err = azlbuildutils.ExecCommandStdout("make",
	toolkitDir,
	"check-manifests")
	return
}

func check_manifests2() (err error) {
	var toolchainScriptsDir = toolkitDir+"/resources/manifests/package"
	err = azlbuildutils.ExecCommandStdout("sh",
		toolchainScriptsDir,
		"update_manifests.sh",
		"aarch64",
		"/home/neha/repos/demo/toolchain_built_rpms_all.tar.gz")
	//toolchain/check_manifests.sh
	return
}

// updateLicenses updates licenses.json file if there are any changes in spec licenses
func updateLicenses() (err error) {
	var script = "license_map.py"
	err = azlbuildutils.ExecCommandStdout("python3",
		scriptsDir,
		script)
	return
}
