// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
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
