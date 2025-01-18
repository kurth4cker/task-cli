// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "testing"

func TestFindNextId(t *testing.T) {
	t.Run("should not duplicate", func(t *testing.T) {
		tasks := []Task{
			{Id: 0},
			{Id: 1},
			{Id: 2},
			{Id: 3},
			{Id: 4},
			{Id: 5},
		}

		got := FindNextId(tasks)
		for _, task := range tasks {
			if got == task.Id {
				t.Errorf("should not give an existing number: %d\n", got)
			}
		}
	})
}
