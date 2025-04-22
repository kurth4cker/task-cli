// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"os"
)

func add(args []string) {
}

func list(args []string) {
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, "list")
	}

	switch args[0] {
	case "add":
		add(args[1:])
	case "list":
		list(args[1:])
	}
}
