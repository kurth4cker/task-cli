// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"encoding/json"
	"os"
)

// Read given file into a Task array
func ReadTasksFile(path string) []Task {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}
		handleError(err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	handleError(err)
	return tasks
}

// Write Task array to given file
func WriteTasksFile(path string, tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", "    ")
	handleError(err)
	data = append(data, byte('\n'))
	err = os.WriteFile(path, data, 0644)
	handleError(err)
}
