// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/kurth4cker/task-cli/internal/task"
)

func TestSet_Add(t *testing.T) {
	freshSet := func(count int) *task.Set {
		set := new(task.Set)
		for i := range count {
			description := fmt.Sprintf("task %d", i+1)
			set.Add(description)
		}
		return set
	}

	t.Run("correct length", func(t *testing.T) {
		set := freshSet(3)
		got := set.Len()
		want := 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("correct and ordered descriptions", func(t *testing.T) {
		set := freshSet(3)
		got := set.Descriptions()
		want := []string{"task 1", "task 2", "task 3"}
		if !slices.Equal(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("correct number of ids", func(t *testing.T) {
		ids := freshSet(3).Ids()
		got := len(ids)
		want := 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("unique ids", func(t *testing.T) {
		isUniq := func(ids []uint) bool {
			for idx, elem := range ids {
				if slices.Contains(ids[idx+1:], elem) {
					return false
				}
			}
			return true
		}

		set := freshSet(3)
		ids := set.Ids()
		if !isUniq(ids) {
			t.Errorf("Ids should be unique, but %+v is not unique", ids)
		}
	})
}

func TestSet_All(t *testing.T) {
	t.Run("should return correct number of elements", func(t *testing.T) {
		tasks := new(task.Set)
		for i := range 15 {
			tasks.Add(fmt.Sprintf("Task %d", i))
		}

		want := 15
		got := 0
		for range tasks.All() {
			got++
		}

		if got != want {
			t.Errorf("got %d elements, want %d", got, want)
		}
	})

	t.Run("should return all descriptions", func(t *testing.T) {
		tasks := new(task.Set)
		tasks.Add("Task 1")
		tasks.Add("Task 2")
		tasks.Add("Task 3")

		want := tasks.Descriptions()
		slices.Sort(want)
		got := []string{}
		for elem := range tasks.All() {
			got = append(got, elem.Description)
		}
		slices.Sort(got)

		if !slices.Equal(got, want) {
			t.Errorf("got descriptions %v, want %v", got, want)
		}
	})
}
