// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"fmt"
	"time"
)

type Task struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`

	// One of "done", "todo", "in-progress"
	Status string `json:"status"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Get string representation of task
func (t Task) String() string {
	return fmt.Sprintf("%d: [%s] %s", t.Id, t.Status, t.Description)
}
