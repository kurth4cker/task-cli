// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

// Set of Task's
type Set struct {
	tasks []Task
}

// Add a new Task with given description and a unique Id
func (s *Set) AddDescription(description string) {
	t := Task{
		Id:          FindNextId(s.tasks),
		Description: description,
	}
	s.tasks = append(s.tasks, t)
}
