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
		file, err := os.OpenFile("database.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open database: %s", err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprintln(file, args[1])
		fmt.Printf("0: %s\n", args[1])
	case "list":
		if len(args) < 2 {
			file, err := os.OpenFile("database.txt", os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot open database: %s", err)
				os.Exit(1)
			}
			defer file.Close()
			_, err = file.WriteTo(os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot print file: %s", err)
				os.Exit(1)
			}
		}
		for _, arg := range args[1:] {
			statuses := []string{"done", "todo", "in-progress"}
			if !slices.Contains(statuses, arg) {
				os.Exit(1)
			}
		}
	}
}
