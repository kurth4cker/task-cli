// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

// Find a new Id for newly added task
func findNextId(tasks []Task) uint {
	var nextId uint
	for _, task := range tasks {
		if task.Id > nextId {
			nextId = task.Id + 1
		}
	}
	return nextId
}
