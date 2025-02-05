// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "tasks.json"

	subcmd := "list"
	flag.Parse()
	if flag.NArg() > 0 {
		subcmd = flag.Arg(0)
	}
	switch subcmd {
	case "list":
		set := readTaskSet(path)
		for task := range set.All() {
			fmt.Println(task)
		}
	case "add":
		if flag.NArg() != 2 {
			fmt.Fprintln(os.Stderr, "wrong usage. provide a description")
			os.Exit(1)
		}
	default:
		log.Fatalf("unknown sub command: %q", subcmd)
	}
}
