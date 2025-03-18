// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main_test

import (
	"log"
	"os"
	"os/exec"
	"strings"
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
	t.Run("should print task names", func(t *testing.T) {
		cases := []struct {
			taskName string
		}{
			{taskName: "Task 1"},
			{taskName: "Task 2"},
		}

		for _, c := range cases {
			output := new(strings.Builder)

			cmd := exec.Command(taskBinary, "add", c.taskName)
			cmd.Stdout = output
			assertCmdRun(t, cmd)

			expected := c.taskName + "\n"
			if output.String() != expected {
				t.Errorf("got %q, want %q", output, expected)
			}
		}
	})

	t.Run("task without task name", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "add")
		assertCmdRun(t, cmd)
		if cmd.ProcessState.Success() {
			t.Error("task-cli succeded, wanted failure")
		}
	})
}

func Test_List(t *testing.T) {
	t.Run("should fail with unknown arguments", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "list", "unknown-status")
		assertCmdRun(t, cmd)
		if cmd.ProcessState.Success() {
			t.Error("'task-cli succeeded, expected to fail")
		}
	})

	t.Run("should not fail without arguments", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "list")
		assertCmdRun(t, cmd)
		if !cmd.ProcessState.Success() {
			t.Error("'task-cli list' failed, expected to succeed")
		}
	})

	t.Run("should not fail with known statuses", func(t *testing.T) {
		statuses := []string{"done", "todo", "in-progress"}
		for _, status := range statuses {
			t.Run(status, func(t *testing.T) {
				cmd := exec.Command(taskBinary, "list", status)
				assertCmdRun(t, cmd)
				if !cmd.ProcessState.Success() {
					t.Errorf("'task-cli list %s' failed, expected to succeed", status)
				}
			})
		}
	})
}

func Test_Raw(t *testing.T) {
	cmd := exec.Command(taskBinary)
	assertCmdRun(t, cmd)

	if !cmd.ProcessState.Success() {
		t.Errorf("task-cli failed, expected to succeed")
	}
}

func assertCmdRun(t testing.TB, cmd *exec.Cmd) {
	t.Helper()
	err := cmd.Run()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return
		default:
			t.Fatalf("cannot run %q: %s", cmd, err)
		}
	}
}
