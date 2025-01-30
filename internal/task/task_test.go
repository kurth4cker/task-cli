// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package task

import (
	"encoding/json"
	"testing"
)

func TestTask_jsonUnmarshal(t *testing.T) {
	t.Run("should work with json.Unmarshal", func(t *testing.T) {
		data := []byte(`
{
    "id": 1,
    "description": "task 1",
    "status": "todo"
}
`)
		var got Task
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatalf("unmarshal failed: %e", err)
		}
		want := Task{
			Id:          1,
			Description: "task 1",
			Status:      "todo",
		}

		if got != want {
			t.Errorf("got %v, want %v, given %q",
				got,
				want,
				string(data))
		}
	})
}
