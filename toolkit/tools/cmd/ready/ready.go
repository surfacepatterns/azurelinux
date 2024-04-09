// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/azlbuild/ready"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("ready", "A command-line interface to ready changes for upstream contribution to azurelinux")
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	err = ready.ReadyChanges()
	if err != nil {
		fmt.Println("Failed to ready changes:\n%w", err)
	}
}
