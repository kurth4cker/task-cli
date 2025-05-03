// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task_test

import (
	"encoding/json"
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
		got := getDescriptions(set)
		want := []string{"task 1", "task 2", "task 3"}
		if !slices.Equal(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("correct number of ids", func(t *testing.T) {
		ids := getIds(freshSet(3))
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
		ids := getIds(set)
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

		want := getDescriptions(tasks);
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

func TestSet_AddElement(t *testing.T) {
	tasks := new(task.Set)
	length := 3
	for i := range length {
		tasks.AddElement(task.Element{
			Id:          uint(i),
			Description: fmt.Sprint("task", i),
		})
	}

	{
		want := length
		got := tasks.Len()
		if got != want {
			t.Errorf("got length %v, want %v", got, want)
		}
	}
}

func TestSet_JSON(t *testing.T) {
	tasks := new(task.Set)
	tasks.Add("task 1")
	tasks.Add("task 2")
	tasks.Add("task 3")

	var want task.Set
	for elem := range tasks.All() {
		want.AddElement(elem)
	}

	var got task.Set
	{
		tmp, err := json.Marshal(&tasks)
		if err != nil {
			t.Fatalf("unexpected marshal error: %s", err)
		}
		if err := json.Unmarshal(tmp, &got); err != nil {
			t.Fatalf("unexpected unmarshal error: %s", err)
		}
	}

	{
		want := slices.Collect(want.All())
		got := slices.Collect(got.All())
		if !unorderedEqual(got, want) {
			t.Errorf("got elements %+v, want %+v", got, want)
		}
	}
}

func unorderedEqual[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	tmp := slices.Clone(s1)
	for _, elem := range s2 {
		idx := slices.Index(tmp, elem)
		if idx == -1 {
			return false
		}
		tmp = remove(tmp, idx)
	}
	return true
}

func remove[S ~[]E, E any](s S, idx int) S {
	if len(s) <= idx {
		return s
	}
	slice := slices.Clone(s[:idx])
	return append(slice, s[idx+1:]...)
}

func getDescriptions(tasks *task.Set) []string {
	descriptions := make([]string, 0, tasks.Len())
	for elem := range tasks.All() {
		descriptions = append(descriptions, elem.Description)
	}
	return descriptions
}

func getIds(s *task.Set) []uint {
	ids := make([]uint, 0, s.Len())
	for elem := range s.All() {
		ids = append(ids, elem.Id)
	}
	return ids
}
