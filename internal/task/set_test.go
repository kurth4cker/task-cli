// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task_test

import (
	"bytes"
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

	t.Run("status defaults to todo", func(t *testing.T) {
		s := new(task.Set)
		s.Add("Task 1")

		got := slices.Collect(s.All())[0].Status
		want := task.Todo
		if got != want {
			t.Errorf("got %v, want %v", got, want)
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
	t.Run("add correct number of elements", func(t *testing.T) {
		tasks := new(task.Set)
		length := 3
		for i := range length {
			tasks.AddElement(task.Element{
				Id:          uint(i),
				Description: fmt.Sprint("task", i),
			})
		}

		want := length
		got := tasks.Len()
		if got != want {
			t.Errorf("got length %v, want %v", got, want)
		}
	})
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

func TestSet_Get(t *testing.T) {
	s := new(task.Set)
	s.AddElement(task.Element{
		Id: uint(1),
		Status: task.Done,
	})
	s.AddElement(task.Element{Id: uint(2)})

	t.Run("found element", func(t *testing.T) {
		elem, ok := s.Get(1)
		if !ok {
			t.Fatalf("cannot found element")
		}

		got := elem.Id
		want := uint(1)
		if got != want {
			t.Errorf("got %v, want %v", got, want);
		}
	})

	t.Run("non-exist element", func(t *testing.T) {
		elem, ok := s.Get(99)
		if ok {
			t.Fatalf("found element")
		}

		got := elem.Id
		want := uint(2)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSet_Mark(t *testing.T) {
	t.Run("should mark element at given id", func(t *testing.T) {
		s := new(task.Set)
		s.AddElement(task.Element{
			Id: uint(1),
			Status: task.Todo,
		})
		s.AddElement(task.Element{
			Id: uint(2),
			Status: task.Todo,
		})

		s.Mark(1, task.Done)

		elem, _ := s.Get(1)
		got := elem.Status
		want := task.Done
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return true on success", func(t *testing.T) {
		s := new(task.Set)
		s.AddElement(task.Element{
			Id: uint(1),
			Status: task.Todo,
		})
		s.AddElement(task.Element{
			Id: uint(2),
			Status: task.Todo,
		})

		ok := s.Mark(1, task.Done)

		got := ok
		want := true
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return false on fail", func(t *testing.T) {
		s := new(task.Set)
		s.AddElement(task.Element{
			Id: uint(1),
			Status: task.Todo,
		})
		s.AddElement(task.Element{
			Id: uint(2),
			Status: task.Todo,
		})

		ok := s.Mark(3, task.InProgress)

		got := ok
		want := false
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSet_ReadFrom(t *testing.T) {
	t.Run("valid json input", func(t *testing.T) {
		s1 := new(task.Set)
		s1.Add("Task 1")
		s1.Add("Task 2")

		data, err := s1.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}

		buf := bytes.NewBuffer(data)
		s2 := new(task.Set)
		if _, err := s2.ReadFrom(buf); err != nil {
			t.Fatal(err)
		}

		got := slices.Collect(s2.All())
		want := slices.Collect(s1.All())
		if !unorderedEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("empty json input", func(t *testing.T) {
		s := new(task.Set)
		_, err := s.ReadFrom(new(bytes.Buffer))
		if err != nil {
			t.Errorf("error occured: %s\n", err)
		}
	})

	t.Run("success read with correct number of bytes", func(t *testing.T) {
		s := new(task.Set)
		buf := bytes.NewBufferString(`[{"Id":0,"Description":"Task 1","Status:"todo"}]`)

		want := int64(buf.Len())
		got, _ := s.ReadFrom(buf)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSet_WriteTo(t *testing.T) {
	s1 := new(task.Set)
	s1.Add("Task 1")
	s1.Add("Task 2")

	buf := new(bytes.Buffer)
	if _, err := s1.WriteTo(buf); err != nil {
		t.Fatal(err)
	}

	s2 := new(task.Set)
	if _, err := s2.ReadFrom(buf); err != nil {
		t.Fatal(err)
	}

	got := slices.Collect(s2.All())
	want := slices.Collect(s1.All())
	if !unorderedEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestSet_Update(t *testing.T) {
	s := new(task.Set)
	s.AddElement(task.Element{
		Description: "wrong description",
		Id: 1,
	})

	correctDescription := "correct description"
	if !s.Update(1, correctDescription) {
		t.Fatal("cannot update task")
	}

	elem, _ := s.Get(1)

	got := elem.Description
	want := correctDescription
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSet_Delete(t *testing.T) {
	s := new(task.Set)
	s.AddElement(task.Element{
		Id: 1,
	})
	s.AddElement(task.Element{
		Id: 2,
	})

	if !s.Delete(1) {
		t.Fatal("cannot delete task")
	}

	got := s.Len()
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSet_Clone(t *testing.T) {
	s1 := new(task.Set)
	s1.Add("Task 1")
	s1.Add("Task 2")

	s2 := s1.Clone()

	got := slices.Collect(s2.All())
	want := slices.Collect(s1.All())
	if !unorderedEqual(got, want) {
		t.Errorf("clone failed")
	}
}

func TestSet_Equal(t *testing.T) {
	t.Run("sets with same descriptions", func(t *testing.T) {
		s1 := new(task.Set)
		s1.Add("Task 1")
		s1.Add("Task 2")

		s2 := new(task.Set)
		s2.Add("Task 1")
		s2.Add("Task 2")

		if !s1.Equal(s2) {
			t.Errorf("should be equal, but not")
		}
	})

	t.Run("sets with different length", func(t *testing.T) {
		s1 := new(task.Set)
		s2 := new(task.Set)
		s1.Add("Task 1")

		if s1.Equal(s2) {
			t.Errorf("should not be equal")
		}
	})

	t.Run("sets with different descriptions", func(t *testing.T) {
		s1 := new(task.Set)
		s2 := new(task.Set)
		s1.Add("Task 1")
		s2.Add("Task 2")

		if s1.Equal(s2) {
			t.Errorf("should not be equal")
		}
	})
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
