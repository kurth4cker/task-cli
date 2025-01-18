// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"slices"
	"testing"
)

func TestFindNextId(t *testing.T) {
	assertGotWant := func(t testing.TB, got, want uint) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertGotNotWant := func(t testing.TB, got, notWant uint) {
		t.Helper()
		if got == notWant {
			t.Errorf("do not want %q but got it", got)
		}
	}

	t.Run("should start from zero", func(t *testing.T) {
		got := findNextId([]Task{})
		var want uint = 0

		assertGotWant(t, got, want)
	})

	t.Run("should not duplicate", func(t *testing.T) {
		tasks := []Task{
			{Id: 0},
			{Id: 1},
			{Id: 2},
			{Id: 3},
			{Id: 4},
			{Id: 5},
		}

		got := findNextId(tasks)
		for _, Task := range tasks {
			assertGotNotWant(t, got, Task.Id)
		}
	})
}

func TestSortedTasks(t *testing.T) {
	compareTaskId := func(a, b Task) int {
		if a.Id < b.Id {
			return -1
		} else if a.Id > b.Id {
			return 1
		} else {
			return 0
		}
	}

	assertSorted := func(t testing.TB, tasks []Task) {
		t.Helper()
		if !slices.IsSortedFunc(tasks, compareTaskId) {
			t.Errorf("task array should be sorted, got %q", tasks)
		}
	}

	t.Run("sorted of sorted", func(t *testing.T) {
		tasks := []Task{
			{Id: 0},
			{Id: 1},
			{Id: 2},
		}
		got := sortedTasks(tasks)
		assertSorted(t, got)
	})

	t.Run("sort normally", func(t *testing.T) {
		tasks := []Task{
			{Id: 3},
			{Id: 2},
			{Id: 0},
		}
		got := sortedTasks(tasks)
		assertSorted(t, got)
	})

	t.Run("do not modify original slice", func(t *testing.T) {
		tasks := []Task{
			{Id: 4},
			{Id: 3},
			{Id: 2},
			{Id: 1},
			{Id: 0},
		}
		oldTasks := slices.Clone(tasks)
		sortedTasks(tasks)
		if slices.CompareFunc(oldTasks, tasks, compareTaskId) != 0 {
			t.Errorf("modified original task array; old %q, new %q",
				oldTasks,
				tasks)
		}
	})
}
