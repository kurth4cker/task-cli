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
	"os"

	"codeberg.org/kurth4cker/task-cli-go/internal/task"
)

func main() {
	tasksPath := "tasks.json"

	subcmd := "list"
	flag.Parse()
	if flag.NArg() > 0 {
		subcmd = flag.Arg(0)
	}
	switch subcmd {
	case "list":
		tasks := task.ReadTasksFile(tasksPath)
		for _, task := range tasks {
			fmt.Println(task)
		}
	case "add":
		if flag.NArg() != 2 {
			fmt.Fprintln(os.Stderr, "wrong usage. provide a description")
			os.Exit(1)
		}
		tasks := task.ReadTasksFile(tasksPath)
		newTask := task.Task{
			Id:          task.FindNextId(tasks),
			Description: flag.Arg(1),
			Status:      "todo",
		}
		tasks = append(tasks, newTask)
		task.WriteTasksFile(tasksPath, tasks)
	default:
		log.Fatalf("unknown sub command: %q", subcmd)
	}
}
