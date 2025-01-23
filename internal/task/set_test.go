// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"bytes"
	"encoding/json"
	"slices"
	"strings"
	"testing"
)

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

	t.Run("should set Status to \"todo\"", func(t *testing.T) {
		var set Set
		set.AddDescription("task 1")

		got := set.tasks[0].Status
		want := "todo"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestSetNewId(t *testing.T) {
	newSetFromIds := func(ids ...uint) (set Set) {
		for _, id := range ids {
			set.tasks = append(set.tasks, Task{Id: id})
		}
		return
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

func TestSetJSONMarshal(t *testing.T) {
	t.Run("should marshal single element Set", func(t *testing.T) {
		var set Set
		set.AddDescription("task 1")
		set.tasks[0].Id = 0
		data, err := json.Marshal(set)
		if err != nil {
			t.Errorf("should marshal to json, failed with %e", err)
		}
		var buffer bytes.Buffer
		if err := json.Compact(&buffer, data); err != nil {
			t.Fatalf("compacting failed: %e", err)
		}

		got := buffer.String()
		want := `[{"id":0,"description":"task 1","status":"todo"}]`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestSetUnmarshalJSON(t *testing.T) {
	t.Run("should unmarshal single element Set", func(t *testing.T) {
		given := []byte(`
[
	{
		"id": 1,
		"description": "task 1",
		"status": "todo"
	}
]
`)
		var set Set
		if err := json.Unmarshal(given, &set); err != nil {
			t.Fatalf("unmarshal failed: %e", err)
		}
		got := set.tasks[0]
		want := Task{
			Id:          1,
			Description: "task 1",
			Status:      "todo",
		}
		if got != want {
			t.Errorf("got %v, want %v, given %s",
				got,
				want,
				string(given))
		}
	})
}

func TestSetReadFrom(t *testing.T) {
	t.Run("should read single element Set", func(t *testing.T) {
		reader := strings.NewReader(`
[
	{
		"id": 1,
		"description": "task 1",
		"status": "todo"
	}
]
`)
		var set Set
		if _, err := set.ReadFrom(reader); err != nil {
			t.Fatalf("read failed: %e", err)
		}

		got := set.tasks[0]
		want := Task{
			Id:          1,
			Description: "task 1",
			Status:      "todo",
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSetWriteTo(t *testing.T) {
	t.Run("should write single element Set", func(t *testing.T) {
		var buffer bytes.Buffer

		var set Set
		set.AddDescription("task 1")
		set.tasks[0].Id = 0

		if _, err := set.WriteTo(&buffer); err != nil {
			t.Fatalf("WriteTo failed: %e", err)
		}
		data := bytes.Clone(buffer.Bytes())
		buffer.Reset()
		if err := json.Compact(&buffer, data); err != nil {
			t.Fatalf("json formatting failed: %e", err)
		}

		got := string(buffer.Bytes())
		want := `[{"id":0,"description":"task 1","status":"todo"}]`
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
