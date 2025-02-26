// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// The package task implements task manipulating type and methods.
//
// You can use it for adding, deleting and tracking a task. Also you can track
// for a task's status by a task id.
//
// Generally you do not need for direct use about Task. You should use Set.
// Which operates all related Task's. You can search and update any Task with
// id.
package task

import "fmt"

type Status string

const (
	Done       = Status("done")
	InProgress = Status("in-progress")
	Todo       = Status("todo")
)

type Task struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`

	// One of "done", "todo", "in-progress"
	Status Status `json:"status"`
}

// Get string representation of task
func (t Task) String() string {
	return fmt.Sprintf("%d: [%s] %s", t.Id, t.Status, t.Description)
}
