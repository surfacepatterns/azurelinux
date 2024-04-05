// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"

	"github.com/microsoft/azurelinux/toolkit/tools/toolinterface/configutils"
)

var (
	// get relevant configs
//	toolkit_dir,_ = configutils.GetBuildConfig("toolkit_root")
	scripts_dir string
)


// ReadyChanges runs various tools to ready changes for contributing to upstream open source repo
// TODO: use a command builder
func ReadyChanges() (err error) {
	fmt.Println("[debug] Ready changes")
	configutils.SetupConfig()
	scripts_dir, _ = configutils.GetBuildConfig("SCRIPTS_DIR")
	fmt.Println("[debug] scripts_dir is ", scripts_dir)

	err = checkManifests()
//	if err != nil {
//		return fmt.Errorf("failed to check manifests:\n%w", err)
//	}

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
	var args = "-a"
	var args2 = "x86_64"
	const ShellToUse = "bash"
	const ShellToUse2 = "-c"

	var script = "toolchain/check_manifests.sh"
	err = execCommands(ShellToUse,
		scripts_dir,
		ShellToUse2,
		script,
		args,
		args2)

	if err != nil {
		err = fmt.Errorf("failed to run (%s):\n%w",script, err)
	}
/*	c = exec.Command("bash -c toolkit/scripts/toolchain/check_manifests.sh -a $arch")

	if err := c.Run(); err != nil {
		err = fmt.Errorf("failed to run check_manifests.sh for arch (%s):\n%w", arch, err)
	}

	arch = "aarch64"
	c = exec.Command("source toolkit/scripts/toolchain/check_manifests.sh")

	if err := c.Run(); err != nil {
		err = fmt.Errorf("failed to run check_manifests.sh for arch (%s):\n%w", arch, err)
	}
	*/
	return
}

// updateLicenses updates licenses.json file if there are any changes in spec licenses
func updateLicenses() (err error) {
	var script = "license_map.py"
	err = execCommands("python3",
		scripts_dir,
		script)

	if err != nil {
		err = fmt.Errorf("failed to run (%s):\n%w",script, err)
	}
	return
}
