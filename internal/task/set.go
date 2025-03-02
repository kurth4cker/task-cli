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
	descriptions := make([]string, len(s.tasks))
	for idx := range len(descriptions) {
		descriptions[idx] = s.tasks[idx].description
	}
	return descriptions
}

func (s *Set) Ids() []uint {
	ids := make([]uint, s.Len())
	for idx := range len(ids) {
		ids[idx] = s.tasks[idx].id
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
			id = task.id + 1
		}
	}
	return id
}

type task struct {
	id          uint
	description string
}
