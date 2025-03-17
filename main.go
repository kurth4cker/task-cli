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
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		os.Exit(0)
	}
	if flag.Arg(0) == "add" {
		if flag.NArg() < 2 {
			os.Exit(1)
		}
		fmt.Println(flag.Arg(1))
	}
}
