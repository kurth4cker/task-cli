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

			if err := cmd.Run(); err != nil {
				t.Fatal("cannot run task-cli", err)
			}

			expected := c.taskName + "\n"
			if output.String() != expected {
				t.Errorf("got %q, want %q", output, expected)
			}
		}
	})

	t.Run("task without task name", func(t *testing.T) {
		cmd := exec.Command(taskBinary, "add")
		err := cmd.Run()
		if !cmd.ProcessState.Exited() {
			t.Fatal("cannot run task-cli", err)
		}
		if cmd.ProcessState.Success() {
			t.Error("task-cli succeded, wanted failure")
		}
	})
}

func Test_Raw(t *testing.T) {
	cmd := exec.Command(taskBinary)
	if err := cmd.Run(); err != nil {
		t.Fatal("cannot run task-cli", err)
	}

	if !cmd.ProcessState.Success() {
		t.Errorf("task-cli failed, expected to succeed")
	}

}
