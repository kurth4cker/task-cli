// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

type Set struct {
	tasks []task
}

func (s *Set) Add(description string) {
	task := task{
		id:          s.newId(),
		description: description,
	}
	s.tasks = append(s.tasks, task)
}

func (s *Set) Descriptions() []string {
	descriptions := make([]string, 0, s.Len())
	for _, task := range s.tasks {
		descriptions = append(descriptions, task.description)
	}
	return descriptions
}

func (s *Set) Ids() []uint {
	ids := make([]uint, 0, s.Len())
	for _, task := range s.tasks {
		ids = append(ids, task.id)
	}
	return ids
}

func (s *Set) Len() int {
	return len(s.tasks)
}

func (s *Set) newId() uint {
	var id uint
	for _, task := range s.tasks {
		if id == task.id {
			id++
		}
	}
	return id
}

type task struct {
	id          uint
	description string
}
