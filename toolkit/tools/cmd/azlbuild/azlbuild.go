// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/buildpackage"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/ready"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app              = kingpin.New("azlbuild", "A command-line interface for azurelinux toolkit")
	build            = app.Command("build", "Build azurelinux")
	setupFunc        = app.Command("setup", "Setup machine")
	readyFunc        = app.Command("ready", "Ready changes to contribute to opensource")

	buildPackageFunc = build.Command("package", "Build package(s)")
	spec             = buildPackageFunc.Flag("spec", "space separated \"\" enclosed name(s) of spec(s) to build").Default("").String()

	buildImageFunc   = build.Command("image", "Build image(s)")
	config           = buildImageFunc.Flag("config", "image config to build").Required().String()
	configDir        = buildImageFunc.Flag("configDir", "directory containing image config").String()
  )

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	// TODO: in each app we should check if the args are correct or not: upfront. first thing to do
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case buildPackageFunc.FullCommand():
			fmt.Println("If this is your first time building Azure Linux, please consider running `setup` to set up your machine")
			fmt.Println("[debug] in build_package ")
			fmt.Println("[debug] spec list is ", *spec)
			err = buildpackage.BuildPackage(*spec)
			if err != nil {
				fmt.Println("[***ERROR***] failed to build package %v", err)
			}
			fmt.Println("Please consider running `ready` before pushing your changes to upstream")
		/*case buildImageFunc.FullCommand():
			fmt.Println("[debug] in image ")
		  	fmt.Println("[debug] config file is ", *config)
		  	err = interfaceutils.BuildImage(*config)
		  	if err != nil {
				fmt.Println("[***ERROR***] failed to build image %v", err)
		}
    	  	fmt.Println("Please consider running `ready` before pushing your changes to upstream")
	*/	case readyFunc.FullCommand():
			fmt.Println("[debug] in ready ")
			err = ready.ReadyChanges()
			if err != nil {
				fmt.Println("[***ERROR***] failed to ready changes %v", err)
			}
		default:
			fmt.Println("[***ERROR***] Invalid option")
	}
}
