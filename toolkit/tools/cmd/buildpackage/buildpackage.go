// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/buildpackage"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app  = kingpin.New("buildpackage", "A command-line interface to build packages in azurelinux")
	spec = app.Flag("spec", "space separated \"\" enclosed name(s) of spec(s) to build").Required().String()
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	fmt.Println("[debug] spec is ", *spec)
	err = buildpackage.BuildPackage(*spec)
	if err != nil {
		fmt.Println("Failed to build package:\n%w", err)
	}
}
