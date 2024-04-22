// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package buildpackage

import (
	"fmt"
	"strings"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/azlbuildutils"
)

var (
	// get relevant configs
	scriptsDir string
	toolchainRpmsDir string
)

// BuildToolchain builds all toolchain packages and packages them into a tar.gz
func BuildToolchain() (err error) {
	DAILY_BUILD_ID, err := runLKG()
	if err != nil {
		err = fmt.Errorf("failed to run LKG script:\n%w", err)
		return err
	}
	azlbuildutils.SetConfig("DAILY_BUILD_ID", DAILY_BUILD_ID)

	err = buildToolchain(DAILY_BUILD_ID)
	if err != nil {
		err = fmt.Errorf("failed to build toolchain:\n%w", err)
		return err
	}

//	err = packageToolchain()
	if err != nil {
		err = fmt.Errorf("failed to package toolchain:\n%w", err)
		return err
	}
	//TODO- set toolchain_ARCHIVE
	return
}

// runLKG updates toolchain RPMs in manifest
func runLKG () (DAILY_BUILD_ID string, err error) {
	scriptsDir, err = azlbuildutils.GetConfig("SCRIPTS_DIR")
	if err != nil {
		err = fmt.Errorf("failed to get config scripts dir:\n%w", err)
		return
	}
	out, err := azlbuildutils.ExecCommandBuffer("sh",
		scriptsDir,
		"setuplkgtoolchain.sh")

	if err != nil {
		return
	}

	// extract DAILY_BUILD_ID from output of setyuplkgtoolchain script
	_, substring, _ := strings.Cut(out, "DAILY_BUILD_ID=")
	_, substring, _ = strings.Cut(substring, "'")
	DAILY_BUILD_ID, _, _ = strings.Cut(substring, "'")
//	fmt.Println("DAILY_BUILD_ID:", DAILY_BUILD_ID)
	return
}

// buildToolchain downloads toolchain RPMs from upstream
func buildToolchain (DAILY_BUILD_ID string) (err error) {
	daily_build_id := "DAILY_BUILD_ID="
	daily_build_id +=DAILY_BUILD_ID
	err = azlbuildutils.ExecCommandStdout("make",
		toolkitDir,
		"toolchain",
		"REBUILD_TOOLS=y",
		"REBUILD_TOOLCHAIN=n",
		daily_build_id)

	if err != nil {
		return
	}
	return
}

// packageToolchain packages built toolchain RPMs into tar.gz
// TODO: implement
func packageToolchain () (err error) {
	// either at TOOLCHAIN_RPMS_DIR or at CACHED_RPMS_DIR
	// store at TOOLCHAIN_ARCHIVE
	// toolchain rpms would be available at build/toolchain_rpms
	toolchainRpmsDir, err = azlbuildutils.GetConfig("TOOLCHAIN_RPMS_DIR")
	if err != nil {
		err = fmt.Errorf("failed to get config toolchainRpmsDir:\n%w", err)
		return
	}
	// copy them and create a tar
	err = azlbuildutils.ExecCommandStdout("tar",
		toolchainRpmsDir,
		"-cvf",
		"",
		"REBUILD_TOOLCHAIN=n")

	if err != nil {
		return
	}
	return
}
