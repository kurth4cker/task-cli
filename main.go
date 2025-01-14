// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"bytes"
	"encoding/json"
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

func main() {
	dataPath := "tasks.json"
	data, err := os.ReadFile(dataPath)
	if err != nil && os.IsNotExist(err) {
		log.Fatalln(err)
	}

	var tasks []task
	json.Unmarshal(data, &tasks)

	// do whatever you want with tasks

	data, err = json.Marshal(tasks)
	var buffer bytes.Buffer
	json.Indent(&buffer, data, "", "    ")
	buffer.WriteRune('\n')
	data = buffer.Bytes()
	err = os.WriteFile(dataPath, data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
