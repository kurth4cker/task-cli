// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package assert_test

import (
	"os/exec"
	"testing"

	"github.com/kurth4cker/task-cli/internal/assert"
)

func TestCmdRun(t *testing.T) {
	cmd := exec.Command("true")
	assert.CmdRun(t, cmd)
}
