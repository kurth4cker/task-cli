// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package assert

import (
	"os/exec"
	"testing"
)

func CmdRun(t testing.TB, cmd *exec.Cmd) {
	t.Helper()
	if err := cmd.Run(); err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return
		default:
			t.Fatalf("cannot run %q: %s", cmd, err)
		}
	}
}
