// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "time"

type Element struct {
	Id          uint
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewElement(description string) Element {
	return Element {
		Description: description,
		Status: Todo,
		CreatedAt: time.Now(),
	}
}

func (e *Element) Touch() {
	e.UpdatedAt = time.Now()
}

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)
