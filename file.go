// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"log"
	"os"

	"codeberg.org/kurth4cker/task-cli-go/internal/task"
)

func newFile(path string) task.Set {
	var set task.Set
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = set.WriteTo(file)
	if err != nil {
		log.Fatal(err)
	}

	return set
}

func readTaskSet(path string) task.Set {
	var set task.Set

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return newFile(path)
		}
		log.Fatal(err)
	}
	defer file.Close()
	_, err = set.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}
	return set
}

func writeTaskSet(path string, set task.Set) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = set.WriteTo(file)
	if err != nil {
		log.Fatal(err)
	}
}
