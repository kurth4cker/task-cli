// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main_test

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/kurth4cker/task-cli/internal/assert"
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
	t.Run("should print task descriptions", func(t *testing.T) {
		cases := []struct {
			taskDescription string
		}{
			{taskDescription: "Task 1"},
			{taskDescription: "Task 2"},
		}

		for _, c := range cases {
			output := new(strings.Builder)

			cmd := exec.Command(taskBinary, "add", c.taskDescription)
			cmd.Stdout = output
			assert.CmdRun(t, cmd)

			expected := c.taskDescription + "\n"
			if output.String() != expected {
				t.Errorf("got %q, want %q", output, expected)
			}
		}
	})

	t.Run("should not fail with task descriptions", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "add", "Task 1")
		assert.CmdRun(t, cmd)
		if !cmd.ProcessState.Success() {
			t.Errorf("%q failed, expected to succeed", cmd)
		}
	})

	t.Run("should fail if no task description", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "add")
		assert.CmdRun(t, cmd)
		if cmd.ProcessState.Success() {
			t.Errorf("%q succeded, expected to fail", cmd)
		}
	})

	// TODO: print added task id
}

func Test_List(t *testing.T) {
	t.Run("should fail with unknown arguments", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "list", "unknown-status")
		assert.CmdRun(t, cmd)
		if cmd.ProcessState.Success() {
			t.Errorf("%q succeeded, expected to fail", cmd)
		}
	})

	t.Run("should not fail without arguments", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "list")
		assert.CmdRun(t, cmd)
		if !cmd.ProcessState.Success() {
			t.Errorf("%q failed, expected to succeed", cmd)
		}
	})

	t.Run("should not fail with known statuses", func(t *testing.T) {
		statuses := []string{"done", "todo", "in-progress"}
		for _, status := range statuses {
			t.Run(status, func(t *testing.T) {
				cmd := exec.Command(taskBinary, "list", status)
				assert.CmdRun(t, cmd)
				if !cmd.ProcessState.Success() {
					t.Errorf("%q failed, expected to succeed", cmd)
				}
			})
		}
	})
}

func Test_Raw(t *testing.T) {
	cmd := exec.Command(taskBinary)
	assert.CmdRun(t, cmd)

	if !cmd.ProcessState.Success() {
		t.Errorf("%q failed, expected to succeed", cmd)
	}
}
