// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

type Element struct {
	Id          uint
	Description string
	Status      Status
}

type Status string

const (
	Todo Status = "todo"
	InProgress = "in-progress"
	Done = "done"
)
