// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/exe"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app    = kingpin.New("buildpackage", "A command-line interface to build images in azurelinux")
	config = app.Flag("config", "path to image config file").Required().String()
)

func main() {
	app.Version(exe.ToolkitVersion)
	var err error
	kingpin.MustParse(app.Parse(os.Args[1:]))
	err = BuildImage(*path)
	if err != nil {
		fmt.Println("Failed to build image:\n%w", err)
	}
}

func BuildImage(path string) (err error) {
	fmt.Println("[debug] Building image from config path", path)
	return
}
