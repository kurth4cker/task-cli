// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"flag"
	"fmt"
	"log"

	"codeberg.org/kurth4cker/task-cli-go/internal/task"
)

func main() {
	tasksPath := "tasks.json"
	tasks := task.ReadTasksFile(tasksPath)

	subcmd := "list"
	flag.Parse()
	if flag.NArg() > 0 {
		subcmd = flag.Arg(0)
	}
	switch subcmd {
	case "list":
		for _, task := range tasks {
			fmt.Println(task)
		}
	default:
		log.Fatalf("unknown sub command: %q", subcmd)
	}

	task.WriteTasksFile(tasksPath, tasks)
}
