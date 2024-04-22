// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package buildpackage

import (
	"fmt"

	packagelist "github.com/microsoft/azurelinux/toolkit/tools/internal/packlist"
	"github.com/microsoft/azurelinux/toolkit/tools/pkg/specreaderutils"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/azlbuildutils"
)

var (
	azlSpecsDirs = [...] string {"SPECS", "SPECS-EXTENDED"}
	// get relevant configs
	toolkitDir string
	projectDir string
)

const (
	// supported build types
	BuildTypeFull     = "full" // build target spec(s) with full specs directory
	BuildTypeIsolated = "isolated" // build target spec(s) in isolation
)

func BuildPackage(spec, buildType string) (err error) {
	// build global configs
//	fmt.Println("[DEBUG] spec is ", spec)
	azlbuildutils.SetupConfig()
	toolkitDir, err = azlbuildutils.GetConfig("toolkit_root")
	if err != nil {
		err = fmt.Errorf("failed to get config toolkit dir:\n%w", err)
		return
	}
//	fmt.Println("[DEBUG] toolkit is ", toolkitDir)
	projectDir, err = azlbuildutils.GetConfig("PROJECT_ROOT")
	if err != nil {
		err = fmt.Errorf("failed to get config project dir:\n%w", err)
		return
	}
//	fmt.Println("[DEBUG] projectDir is ", projectDir)

//	fmt.Println("[DEBUG] Building packages: specs are (%s)", spec)

	// check specs exist
	specsDir, err := validateSpecExistance(spec)
	if err != nil {
		err = fmt.Errorf("failed to validate specs:\n%w", err)
		return
	}

	// TODO: set sepcs dir in config

	// any other checks

	// build toolchain if required

	err = BuildToolchain()
	if err != nil {
		err = fmt.Errorf("failed to build toolchain:\n%w", err)
		return
	}
	// put toolchain rpms into toolchain_archive and use it

	// set extra configs

	// show dependency graph - use graphanalytics tool

	// build package
	err = buildSpecs(spec, specsDir, buildType)
	if err != nil {
		err = fmt.Errorf("failed to build specs:\n%w", err)
		return err
	}

	// show output

	return
}

// validateSpecExistance checks if each spec in specList exists
// If the spec exists, it assigns it the correct specsDir
func validateSpecExistance(specList string) (specsDir string, err error) {
//	fmt.Println("[DEBUG] Checking if spec exists for (%s)", specList)
	specMap, err := packagelist.ParsePackageList(specList)
	if err != nil {
		err = fmt.Errorf("failed to parse package list:\n%w", err)
		return
	}

	// TODO: currently, we have a limitation that all specs to be built must be present in the same specsDir
	// TODO: return error only if spec is not found in any specsDir
	// TODO: read specs dir from config if it has been redefined by user in config json
	for _, specsDir := range azlSpecsDirs {
		_, err := specreaderutils.FindSpecFiles(projectDir+specsDir, specMap)
		if err != nil {
			err = fmt.Errorf("failed to FindSpecFiles:\n%w", err)
			return "", err
		} else {
//			fmt.Println("[DEBUG] done with specreader, returned specFiles (%s)", specFiles)
			return projectDir+specsDir, nil
		}
	}
//	fmt.Println("[DEBUG] done with specreader")
	return
}

// buildSpecs builds specs in specsDir either in full or isolation
func buildSpecs (specs, specsDir, buildType string) (err error) {
	// TODO: use a command builder
	// TODO: some of these arguments can be removed if/when tools start reading directly from config
	// TODO: build in full by default
	// TODO: add toolchain_archive to make

	switch buildType {
	case BuildTypeFull:
		return buildSpecsFull(specs, specsDir)
	case BuildTypeIsolated:
		return buildSpecsIsolated(specs, specsDir)
	default:
		err = fmt.Errorf("incorrect buildType for building packages")
		return
	}
	return
}

// buildSpecsFull builds specs in specsDir in full
func buildSpecsFull(specs, specsDir string) (err error) {
	specs_dir := "SPECS_DIR="
	specs_dir += specsDir
	package_build_list := "PACKAGE_BUILD_LIST="
	package_build_list +=specs
	package_rebuild_list := "PACKAGE_REBUILD_LIST="
	package_rebuild_list +=specs

	err = azlbuildutils.ExecCommandStdout("make",
		toolkitDir,
		"build-packages",
		"REBUILD_TOOLS=y",
		package_build_list,
		package_rebuild_list)
		//specs_dir)

	if err != nil {
		return
	}
	return
}

// buildSpecsIsolated builds specs in specsDir in isolation
func buildSpecsIsolated(specs, specsDir string) (err error) {
	specs_dir := "SPECS_DIR="
	specs_dir += specsDir
	srpm_pack_list := "SRPM_PACK_LIST="
	srpm_pack_list +=specs

	err = azlbuildutils.ExecCommandStdout("make",
		toolkitDir,
		"build-packages",
		"REBUILD_TOOLS=y",
		srpm_pack_list,
		specs_dir)

	if err != nil {
		return
	}
	return
}
