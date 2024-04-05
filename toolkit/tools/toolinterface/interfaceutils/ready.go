// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"

	"github.com/microsoft/azurelinux/toolkit/tools/toolinterface/configutils"
)

var (
	// get relevant configs
	scripts_dir string
	toolkitDir string
)


// ReadyChanges runs various tools to ready changes for contributing to upstream open source repo
// TODO: use a command builder
func ReadyChanges() (err error) {
	fmt.Println("[debug] Ready changes")
	configutils.SetupConfig()
	scripts_dir, _ = configutils.GetBuildConfig("SCRIPTS_DIR")
	toolkitDir, _ = configutils.GetBuildConfig("toolkit_root")
	fmt.Println("[debug] scripts_dir is ", scripts_dir)

	err = checkManifests()
	if err != nil {
		return fmt.Errorf("failed to check manifests:\n%w", err)
	}

	err = updateLicenses()
	if err != nil {
		return fmt.Errorf("failed to update licenses:\n%w", err)
	}
	return
}

func specLint() (err error) {
	return
}

// checkManifest runs check_manifests script to spot updates required in manifest files
func checkManifests() (err error) {
	err = execCommands("make",
	toolkitDir,
	"check-manifests")
	return
}

// updateLicenses updates licenses.json file if there are any changes in spec licenses
func updateLicenses() (err error) {
	var script = "license_map.py"
	err = execCommands("python3",
		scripts_dir,
		script)
	return
}
