// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "slices"

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

// Return sorted copy of given task slice. Tasks are compared by Id fields.
func sortedTasks(tasks []Task) []Task {
	compareTaskId := func(a, b Task) int {
		if a.Id < b.Id {
			return -1
		} else if a.Id > b.Id {
			return 1
		} else {
			return 0
		}
	}

	return slices.SortedFunc(slices.Values(tasks), compareTaskId)
}
