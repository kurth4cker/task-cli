// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kurth4cker/task-cli/internal/task"
)

func TestNewElement(t *testing.T) {
	t.Run("properly setup description", func(t *testing.T) {
		description := "Task 1"
		elem := task.NewElement(description)

		got := elem.Description
		want := description
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("status defaults to todo", func(t *testing.T) {
		elem := task.NewElement("Task 1")

		got := elem.Status
		want := task.Todo
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("create with fresh time", func(t *testing.T) {
		elem := task.NewElement("Task 1")

		got := elem.CreatedAt
		want := time.Now()

		if got.Sub(want).Abs() > time.Second {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestElement_JSON(t *testing.T) {
	element := task.Element{
		Id:          0,
		Description: "Task 1",
	}

	var got task.Element
	want := element
	{
		tmp, err := json.Marshal(&element)
		if err != nil {
			t.Fatalf("marshal, unexpected error: %s", err)
		}
		err = json.Unmarshal(tmp, &got)
		if err != nil {
			t.Fatalf("unmarshal, unexpected error: %s", err)
		}
	}

	if got != want {
		t.Errorf("got element %v, want %v", got, want)
	}
}

func TestElement_Touch(t *testing.T) {
	elem := task.Element {
		Description: "Task 1",
		Status: task.Todo,
	}

	elem.Touch()

	got := elem.UpdatedAt
	want := time.Now()

	if got.Sub(want).Abs() > time.Second {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStatus_String(t *testing.T) {
	cases := []struct{
		given task.Status
		want string
	}{
		{given: task.Done, want: "done"},
		{given: task.InProgress, want: "in-progress"},
		{given: task.Todo, want: "todo"},
	}

	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			want := c.want
			got := string(c.given)
			if got != want {
				t.Errorf("got %v, want %v, given %v", got, want, c.given)
			}
		})
	}
}
