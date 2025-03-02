// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package data

import "os"

// Ensure given path is exist.
func EnsureFile(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return
		}
		file.Close()
		return
	}
}
