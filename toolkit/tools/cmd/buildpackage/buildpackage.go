// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/buildpackage"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("buildpackage", "A command-line interface to build packages in azurelinux")
	spec     = app.Flag("spec", "space separated \"\" enclosed name(s) of spec(s) to build").Required().String()
	logFlags = exe.SetupLogFlags(app)
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	logger.InitBestEffort(logFlags)
	logger.Log.Debugf("spec is ", *spec)
	err = buildpackage.BuildPackage(*spec)
	if err != nil {
		logger.Log.Fatalf("Failed to build package:\n%v", err)
	}
}
