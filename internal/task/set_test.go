// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import "testing"

func TestSetAddDescription(t *testing.T) {
	t.Run("should add to empty Set", func(t *testing.T) {
		var set Set
		set.AddDescription("task 1")
		set.AddDescription("task 2")

		want := 2
		got := len(set.tasks)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("should not change description", func(t *testing.T) {
		var set Set
		set.AddDescription("task 1")

		want := "task 1"
		got := set.tasks[0].Description
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("should add with unique ids", func(t *testing.T) {
		toIdArray := func(s Set) []uint {
			var ids []uint
			for _, task := range s.tasks {
				ids = append(ids, task.Id)
			}
			return ids
		}

		var set Set
		set.AddDescription("task 1")
		set.AddDescription("task 2")
		set.AddDescription("task 3")

		ids := toIdArray(set)
		for k, id := range ids {
			for x := k + 1; x < len(ids); x++ {
				if ids[x] == id {
					t.Errorf("duplicate ids: %v", ids)
				}
			}
		}
	})
}
