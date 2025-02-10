// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"slices"
	"testing"
)

func TestSetAddDescription(t *testing.T) {
	t.Run("should add correct number of elements to empty Set",
		func(t *testing.T) {
			var set Set
			set.AddDescription("task 1")
			set.AddDescription("task 2")

			want := 2
			got := len(set)
			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})

	t.Run("should not change description", func(t *testing.T) {
		var set Set
		set.AddDescription("task 1")

		for task := range set.All() {
			want := "task 1"
			got := task.Description
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})

	t.Run("should add with unique ids", func(t *testing.T) {
		toIdSlice := func(s Set) []uint {
			var ids []uint
			for task := range s.All() {
				ids = append(ids, task.Id)
			}
			return ids
		}

		var set Set
		set.AddDescription("task 1")
		set.AddDescription("task 2")
		set.AddDescription("task 3")

		ids := toIdSlice(set)
		for k, id := range ids {
			for x := k + 1; x < len(ids); x++ {
				if ids[x] == id {
					t.Errorf("duplicate ids: %v", ids)
				}
			}
		}
	})

	t.Run("should set Status to \"todo\"", func(t *testing.T) {
		var set Set
		set.AddDescription("task 1")

		for task := range set.All() {
			got := task.Status
			want := "todo"

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
	})
}

func TestSetNewId(t *testing.T) {
	newSetFromIds := func(ids ...uint) Set {
		var set Set
		for _, id := range ids {
			set = append(set, Task{Id: id})
		}
		return set
	}

	t.Run("should be unique", func(t *testing.T) {
		ids := []uint{0, 1, 2, 3, 4}
		set := newSetFromIds(ids...)

		newId := set.newId()

		if slices.Contains(ids, newId) {
			t.Errorf("not unique: given %v, got %d", ids, newId)
		}
	})
}

func TestSetAll(t *testing.T) {
	t.Run("should iterate over all Tasks", func(t *testing.T) {
		var set Set
		set.AddDescription("task 0")
		set.AddDescription("task 1")

		var got []string
		for task := range set.All() {
			got = append(got, task.Description)
		}
		want := []string{"task 0", "task 1"}
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
