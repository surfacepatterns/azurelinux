// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("buildpackage", "A command-line interface to build images in azurelinux")
	config   = app.Flag("config", "path to image config file").Required().String()
	logFlags = exe.SetupLogFlags(app)
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	logger.InitBestEffort(logFlags)
	err = BuildImage(*config)
	if err != nil {
		logger.Log.Fatalf("Failed to build image:\n%w", err)
	}
}

func BuildImage(config string) (err error) {
	logger.Log.Debugf("Building image from config (%s)", config)
	return
}
