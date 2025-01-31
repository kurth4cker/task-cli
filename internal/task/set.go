// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"encoding/json"
	"io"
	"iter"
)

// Set of Task's
type Set struct {
	tasks []Task
}

// Add given task to the Set.
func (s *Set) Add(t Task) {
	s.tasks = append(s.tasks, t)
}

// Add a new Task with given description and a unique Id.
func (s *Set) AddDescription(description string) {
	t := Task{
		Id:          s.newId(),
		Description: description,
		Status:      "todo",
	}
	s.tasks = append(s.tasks, t)
}

func (s Set) All() iter.Seq[Task] {
	return func(yield func(Task) bool) {
		for _, v := range s.tasks {
			if !yield(v) {
				return
			}
		}
	}
}

// Returns JSON representation of Set as JSON Array.
//
// Each element is a [Task] object.
func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.tasks)
}

// Read from given Reader.
//
// Reader should contain json encoded Task Set data. ReadFrom decodes and
// appends all elements into Set.
func (s *Set) ReadFrom(r io.Reader) (int64, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return int64(len(data)), err
	}
	err = json.Unmarshal(data, s)
	return int64(len(data)), err
}

// Unmarshal a JSON encoded Task array into Set.
func (s *Set) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.tasks)
}

// Write to given Writer as indented json encoded data.
func (s *Set) WriteTo(w io.Writer) (int64, error) {
	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return int64(len(data)), err
	}
	n, err := w.Write(data)
	return int64(n), err
}

// Generate a new Id which is not found in tasks
//
// This Id typically used for adding a new Task to Set.
func (s *Set) newId() uint {
	var maxId uint
	for _, task := range s.tasks {
		if maxId < task.Id {
			maxId = task.Id
		}
	}
	return maxId + 1
}
