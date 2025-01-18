// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

// Find a new Id for newly added task
func FindNextId(tasks []Task) uint {
	var maxId uint
	for _, task := range tasks {
		if maxId < task.Id {
			maxId = task.Id
		}
	}
	return maxId + 1
}
