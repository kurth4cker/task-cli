// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main_test

import (
	"os"
	"os/exec"
	"testing"
)

func Test_Add(t *testing.T) {
	binaryPath := "./task-cli"

	buildCmd := exec.Command("go", "build", "-o", binaryPath, ".")
	if err := buildCmd.Run(); err != nil {
		t.Fatal("cannot compile task-cli", err)
	}
	defer os.Remove(binaryPath)

	taskName := "Test Task"
	cmd := exec.Command(binaryPath, "add", taskName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("cannout run task-cli", err)
	}

	expected := taskName + "\n"
	if string(output) != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}
