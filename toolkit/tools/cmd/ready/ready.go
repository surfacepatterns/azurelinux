// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/toolinterface/toolinterfaceutils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("ready", "A command-line interface to ready changes for upstream contribution to azurelinux")
)

var (
	// get relevant configs
	toolkitDir string
	scriptsDir string
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	err = ReadyChanges()
	if err != nil {
		fmt.Println("Failed to ready changes:\n%w", err)
	}
}

// ReadyChanges runs various tools to ready changes for contributing to upstream open source repo
// TODO: use a command builder
func ReadyChanges() (err error) {
	fmt.Println("[debug] Ready changes")
	toolinterfaceutils.SetupConfig()
	scriptsDir, _ = toolinterfaceutils.GetBuildConfig("SCRIPTS_DIR")
	toolkitDir, _ = toolinterfaceutils.GetBuildConfig("toolkit_root")
	fmt.Println("[debug] scripts_dir is ", scriptsDir)

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
	err = toolinterfaceutils.ExecCommands("make",
	toolkitDir,
	"check-manifests")
	return
}

// updateLicenses updates licenses.json file if there are any changes in spec licenses
func updateLicenses() (err error) {
	var script = "license_map.py"
	err = toolinterfaceutils.ExecCommands("python3",
		scriptsDir,
		script)
	return
}
