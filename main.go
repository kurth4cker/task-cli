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
		set, err := task.NewSetFile(tasksPath)
		if err != nil {
			if os.IsNotExist(err) {
				set = &task.Set{}
			} else {
				log.Fatalf("cannoct read file: %s", err)
			}
		}
		for task := range set.All() {
			fmt.Println(task)
		}
		set.WriteFile(tasksPath)
	case "add":
		if flag.NArg() != 2 {
			fmt.Fprintln(os.Stderr, "wrong usage. provide a description")
			os.Exit(1)
		}
		set, err := task.NewSetFile(tasksPath)
		if err != nil {
			if os.IsNotExist(err) {
				set = &task.Set{}
			} else {
				log.Fatalf("error: %e", err)
			}
		}
		set.AddDescription(flag.Arg(1))
		set.WriteFile(tasksPath)
	default:
		log.Fatalf("unknown sub command: %q", subcmd)
	}
}
