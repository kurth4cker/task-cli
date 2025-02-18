// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"slices"
	"testing"

	"codeberg.org/kurth4cker/task-cli/internal/task"
)

func TestSet_AddDescription(t *testing.T) {
	want := []string{
		"task 1",
		"task 2",
	}

	var set task.Set
	for _, description := range want {
		set.AddDescription(description)
	}

	var got []string
	for task := range set.All() {
		got = append(got, task.Description)
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSet_All(t *testing.T) {
	t.Run("empty task set", func(t *testing.T) {
		want := 0

		var set task.Set
		got := len(slices.Collect(set.All()))

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
