// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package data_test

import (
	"os"
	"testing"

	"codeberg.org/kurth4cker/task-cli/internal/data"
)

func TestEnsureFile(t *testing.T) {
	t.Run("file exist", func(t *testing.T) {
		testFile := "testdata/exist.json"
		data.EnsureFile(testFile)

		info, err := os.Stat(testFile)
		if err != nil {
			t.Fatal(err)
		}
		if !info.Mode().IsRegular() {
			t.Errorf("%s should be regular file but it is not", testFile)
		}
	})

	t.Run("file not exist", func(t *testing.T) {
		testFile := "testdata/none.json"
		data.EnsureFile(testFile)

		info, err := os.Stat(testFile)
		if err != nil {
			t.Fatal(err)
		}
		if !info.Mode().IsRegular() {
			t.Errorf("%s should be regular file but it is not", testFile)
		}

		// cleanup after tests
		os.Remove(testFile)
	})
}
