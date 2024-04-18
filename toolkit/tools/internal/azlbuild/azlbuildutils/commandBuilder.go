// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azlbuildutils

import (
	"io"
	"fmt"
    "os"
	"os/exec"
)

// buildStatus checks if spec needs to be rebuilt
// reasons for rebuild:
//					- spec has never been built
//					- there is a change in spec from last build
//					- there is a change in toolchain manifest
//					- user wants to rebuild spec
func buildStatus() (err error) {
	//fmt.Println("[DEBUG] in buildStatus")
	return

}

// buildStatusToolchain checks if toolchain should be (re)built
// returns true if
//					- toolchain rpms have never been built
//					- there is a change in toolchain spec(s)
//					- there is a change in toolchain manifest
//					- user wants to rebuild toolchain
func buildStatusToolchain() (rebuildOpt bool, err error) {
	err = buildStatus()
	// added if change in manifest files
	return false, nil
}

// ExecCommandStdout executes the application from the directory with the given arguments, and redirects output to stdout
func ExecCommandStdout(app, dir string, args ...string) (err error) {
	cmd := exec.Command(app, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = io.MultiWriter(os.Stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr)

	err = cmd.Start()
	if err != nil {
		err = fmt.Errorf("failed to exec cmd.Start():\n%w", err)
        return
	}
	err = cmd.Wait()
	if err != nil {
		err = fmt.Errorf("failed to exec cmd.Run():\n%w", err)
        return
	}
    return
}

// ExecCommandBuffer executes the application from the directory with the given arguments, and stores output to out buffer
func ExecCommandBuffer(app, dir string, args ...string) (out string, err error) {
	cmd := exec.Command(app, args...)
	if dir != "" {
		cmd.Dir = dir
	}

	outBuff, err := cmd.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("failed to exec cmd.Run():\n%w", err)
        return
	}
	out = string(outBuff)
	return
}
