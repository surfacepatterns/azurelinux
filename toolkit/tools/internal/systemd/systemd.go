// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package systemd

import (
	"fmt"
	"strings"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/safechroot"
	"github.com/microsoft/azurelinux/toolkit/tools/internal/shell"
)

// IsServiceEnabled checks if a service is enabled or disabled.
func IsServiceEnabled(name string, imageChroot safechroot.ChrootInterface) (bool, error) {
	serviceEnabled := true
	err := imageChroot.UnsafeRun(func() error {
		stdout, stderr, err := shell.Execute("systemctl", "is-enabled", name)

		// `systemctl is-enabled` returns:
		//   enabled:  Exit code = 0, stdout = "enabled"
		//   disabled: Exit code = 1, stdout = "disabled"
		//   error:    Exit code = 1, stdout = ""
		if err != nil {
			if strings.TrimSpace(stdout) != "disabled" {
				return fmt.Errorf("%s\n%w", strings.TrimSpace(stderr), err)
			}

			serviceEnabled = false
		}

		return nil
	})
	if err != nil {
		return false, fmt.Errorf("failed to check if (%s) service is enabled:\n%w", name, err)
	}

	return serviceEnabled, nil
}
