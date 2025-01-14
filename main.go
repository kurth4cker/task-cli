// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
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

// If there is an error, print it
func maybe(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	data, err := json.Marshal(tasks)
	maybe(err)

	var buffer bytes.Buffer
	json.Indent(&buffer, data, "", "    ")
	buffer.WriteRune('\n')
	data = buffer.Bytes()
	err = os.WriteFile(path, data, 0644)
	maybe(err)
}

func main() {
	tasksPath := "tasks.json"
	tasks := readTasksFile(tasksPath)

	var subcmd string
	if flag.NArg() == 0 {
		subcmd = "list"
	}
	switch subcmd {
	case "list":
		for _, task := range tasks {
			fmt.Println(task)
		}
	}

	writeTasksFile(tasksPath, tasks)
}
