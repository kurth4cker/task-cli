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
	if flag.NArg() < 1 {
		os.Exit(0)
	}

	switch flag.Arg(0) {
	case "add":
		if flag.NArg() < 2 {
			os.Exit(1)
		}
		fmt.Println(flag.Arg(1))
	case "list":
		if flag.NArg() < 2 {
			os.Exit(0)
		}
		args := flag.Args()[1:]
		statuses := []string{"done", "todo", "in-progress"}
		for _, arg := range args {
			if !slices.Contains(statuses, arg) {
				os.Exit(1)
			}
		}
	}
}
