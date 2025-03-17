// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main_test

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

const taskBinary = "./task-cli"

func TestMain(m *testing.M) {
	buildCmd := exec.Command("go", "build", "-o", taskBinary, ".")
	if err := buildCmd.Run(); err != nil {
		log.Fatal("cannot compile task-cli", err)
	}
	code := m.Run()
	os.Remove(taskBinary)

	os.Exit(code)
}

func Test_Add(t *testing.T) {
	cases := []struct {
		taskName string
	}{
		{taskName: "Task 1"},
		{taskName: "Task 2"},
	}

	for _, c := range cases {
		cmd := exec.Command(taskBinary, "add", c.taskName)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal("cannout run task-cli", err)
		}

		expected := c.taskName + "\n"
		if string(output) != expected {
			t.Errorf("got %q, want %q", output, expected)
		}
	}
}

func Test_Raw(t *testing.T) {
	cmd := exec.Command(taskBinary)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal("cannot run task-cli", err)
	}

	expected := ""
	if string(output) != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}
