// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "iter"

// Set of Task's
type Set []Task

// Add a new Task with given description and a unique Id.
func (s *Set) AddDescription(description string) {
	*s = append(*s, Task{
		Id:          s.newId(),
		Description: description,
		Status:      "todo",
	})
}

// Return an iterator which iterates over all tasks.
func (s *Set) All() iter.Seq[Task] {
	return func(yield func(Task) bool) {
		for _, v := range *s {
			if !yield(v) {
				return
			}
		}
	}
}

// Generate a new Id which is not found in tasks
//
// This Id typically used for adding a new Task to Set.
func (s *Set) newId() uint {
	var maxId uint
	for _, task := range *s {
		if maxId < task.Id {
			maxId = task.Id
		}
	}
	return maxId + 1
}
