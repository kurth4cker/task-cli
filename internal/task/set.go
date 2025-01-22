// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "encoding/json"

// Set of Task's
type Set struct {
	tasks []Task
}

// Add a new Task with given description and a unique Id
func (s *Set) AddDescription(description string) {
	t := Task{
		Id:          s.newId(),
		Description: description,
		Status:      "todo",
	}
	s.tasks = append(s.tasks, t)
}

// Returns JSON representation of Set as JSON Array.
//
// Each element is a [Task] object.
func (s Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.tasks)
}

func (s *Set) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.tasks)
}

// Generate a new Id which is not found in tasks
//
// This Id typically used for adding a new Task to Set.
func (s Set) newId() uint {
	var maxId uint
	for _, task := range s.tasks {
		if maxId < task.Id {
			maxId = task.Id
		}
	}
	return maxId + 1
}
