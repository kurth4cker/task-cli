// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kurth4cker/task-cli/internal/task"
)

const (
	taskFileName = "tasks.json"
)

func add(args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Wron number of arguments. Usage:")
		fmt.Fprintf(os.Stderr, "    add <task description>\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(taskFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1);
	}
	defer f.Close()

	// read data
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// unmarshal data into set
	set := new(task.Set)
	if len(data) != 0 {
		err = set.UnmarshalJSON(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot parse %s: %s\n", taskFileName, err)
			os.Exit(1)
		}
	}

	// add new task and marshal again
	set.Add(args[0])
	data, err = set.MarshalJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create json output: %s\n", err)
		os.Exit(1)
	}

	// write data
	_, err = f.WriteAt(data, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot write to file: %s\n", err)
		os.Exit(1)
	}
}

func list(args []string) {
	// TODO(#6): implement list
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, "list")
	}

	switch args[0] {
	case "add":
		add(args[1:])
	case "list":
		list(args[1:])
	}
}
