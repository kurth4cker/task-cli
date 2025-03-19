// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		args = append(args, "list")
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			os.Exit(1)
		}
		fmt.Printf("0: %s", args[1])
	case "list":
		if len(args) < 2 {
			os.Exit(0)
		}
		for _, arg := range args[1:] {
			statuses := []string{"done", "todo", "in-progress"}
			if !slices.Contains(statuses, arg) {
				os.Exit(1)
			}
		}
	}
}
