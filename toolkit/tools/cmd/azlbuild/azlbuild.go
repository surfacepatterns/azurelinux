// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/buildpackage"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/ready"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app              = kingpin.New("azlbuild", "A command-line interface for azurelinux toolkit")
	build            = app.Command("build", "Build azurelinux")
	setupFunc        = app.Command("setup", "Setup machine")
	readyFunc        = app.Command("ready", "Ready changes to contribute to opensource")

	buildPackageFunc = build.Command("package", "Build package(s)")
	spec             = buildPackageFunc.Flag("spec", "space separated \"\" enclosed name(s) of spec(s) to build").Default("").String()
	buildType        = buildPackageFunc.Flag("buildType", "build spec(s) with full AZL or isolated. Supported: full, isolated").Required().Enum("full", "isolated")

	buildImageFunc   = build.Command("image", "Build image(s)")
	config           = buildImageFunc.Flag("config", "image config to build").Required().String()
	configDir        = buildImageFunc.Flag("configDir", "directory containing image config").String()

	logFlags         = exe.SetupLogFlags(app)
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	logger.InitBestEffort(logFlags)

	// TODO: in each app we should check if the args are correct or not: upfront. first thing to do
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case buildPackageFunc.FullCommand():
			logger.Log.Infof("If this is your first time building Azure Linux, please consider running `setup` to set up your machine")
			logger.Log.Debugf("in build_package ")
			logger.Log.Debugf("spec list:(%s)", *spec)
			if *buildType == "" {
				kingpin.Fatalf("buildType must be specified. Supported: full, isolated")
			}
			err = buildpackage.BuildPackage(*spec, *buildType)
			if err != nil {
				logger.Log.Fatalf("failed to build package %v", err)
			}
			logger.Log.Infof("Please consider running `ready` before pushing your changes to upstream")
		/*case buildImageFunc.FullCommand():
			logger.Log.Debugf("in image ")
			logger.Log.Debugf("config file is ", *config)
		  	err = interfaceutils.BuildImage(*config)
		  	if err != nil {
				logger.Log.Fatalf("failed to build image %v", err)
		}
			logger.Log.Infof("Please consider running `ready` before pushing your changes to upstream")
	*/	case readyFunc.FullCommand():
			logger.Log.Debugf("in ready ")
			err = ready.ReadyChanges()
			if err != nil {
				logger.Log.Fatalf("failed to ready changes %v", err)
			}
		default:
			logger.Log.Fatalf("Invalid option")
	}
}
