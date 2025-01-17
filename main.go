// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

type task struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`

	// One of "done", "todo", "in-progress"
	Status string `json:"status"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Get string representation of task
func (t task) String() string {
	return fmt.Sprintf("%d: [%s] %s", t.Id, t.Status, t.Description)
}

func findNextId(tasks []task) uint {
	var nextId uint
	for _, task := range tasks {
		if task.Id > nextId {
			nextId = task.Id + 1
		}
	}
	return nextId
}

// If there is an error, print it
func maybe(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Return sorted copy of given task slice
func sortedTasks(tasks []task) []task {
	compareTaskId := func(a, b task) int {
		if a.Id < b.Id {
			return -1
		} else if a.Id > b.Id {
			return 1
		} else {
			return 0
		}
	}

	return slices.SortedFunc(slices.Values(tasks), compareTaskId)
}

// Read given file
func readTasksFile(path string) []task {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []task{}
		}
		maybe(err)
	}

	var tasks []task
	err = json.Unmarshal(data, &tasks)
	maybe(err)
	return tasks
}

func writeTasksFile(path string, tasks []task) {
	data, err := json.MarshalIndent(tasks, "", "    ")
	maybe(err)
	data = append(data, byte('\n'))
	err = os.WriteFile(path, data, 0644)
	maybe(err)
}

func main() {
	tasksPath := "tasks.json"
	tasks := readTasksFile(tasksPath)

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

	writeTasksFile(tasksPath, tasks)
}
