// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/ready"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("ready", "A command-line interface to ready changes for upstream contribution to azurelinux")
	logFlags = exe.SetupLogFlags(app)
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	logger.InitBestEffort(logFlags)
	err = ready.ReadyChanges()
	if err != nil {
		logger.Log.Fatalf("Failed to ready changes:\n%v", err)
	}
}
