// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Task management CLI application
//
// Created for a project idea taken from roadmap.sh
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kurth4cker/task-cli/internal/task"
)

const (
	taskFileName = "tasks.json"
)

func add(args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Wron number of arguments. Usage:")
		fmt.Fprintf(os.Stderr, "    add <task description>\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(taskFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1);
	}
	defer f.Close()

	set := new(task.Set)
	if _, err := set.ReadFrom(f); err != nil {
		fmt.Fprintf(os.Stderr, "cannot read and parse %s: %s\n", taskFileName, err)
		os.Exit(1)
	}

	set.Add(args[0])

	f.Seek(0, 0)
	n, err := set.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create and write json output: %s\n", err)
		os.Exit(1)
	}
	f.Truncate(n)
}

func list(args []string) {
	if len(args) > 2 {
		fmt.Fprintf(os.Stderr, "invalid number of arguments. usage:\n")
		fmt.Fprintf(os.Stderr, "    list [status]\n")
		os.Exit(1)
	}
	var status task.Status
	listAll := true
	if len(args) == 1 {
		switch (args[0]) {
		case "todo":
			status = task.Todo
		case "in-progress":
			status = task.InProgress
		case "done":
			status = task.Done
		}
		listAll = false
	}

	data, err := os.ReadFile(taskFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	set := new(task.Set)
	err = set.UnmarshalJSON(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse %s: %s\n", taskFileName, err)
		os.Exit(1)
	}

	if listAll {
		for elem := range set.All() {
			fmt.Printf("%v, %v: %s\n", elem.Status, elem.Id, elem.Description)
		}
	} else {
		// TODO(#26): add filter functionality to task.Set
		for elem := range set.All() {
			if elem.Status == status {
				fmt.Printf("%v, %v: %s\n", elem.Status, elem.Id, elem.Description)
			}
		}
	}
}

func mark(status task.Status, args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "invalid usage. please specify only one Id\n")
		os.Exit(1)
	}

	var id uint
	if id64, err := strconv.ParseUint(args[0], 10, 32); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse id: %s\n", err)
		os.Exit(1)
	} else {
		id = uint(id64)
	}

	f, err := os.OpenFile(taskFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1);
	}
	defer f.Close()

	set := new(task.Set)
	if _, err := set.ReadFrom(f); err != nil {
		fmt.Fprintf(os.Stderr, "cannot read and parse %s: %s\n", taskFileName, err)
		os.Exit(1)
	}

	// TODO(#24): check for errors
	set.Mark(id, status)

	f.Seek(0, 0)
	n, err := set.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create and write json output: %s\n", err)
		os.Exit(1)
	}
	f.Truncate(n)
}

func update(args []string) {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "invalid number of arguments. usage:\n")
		fmt.Fprintf(os.Stderr, "    update <id> <new description>\n")
		os.Exit(1)
	}

	id := parseId(args[0])

	f, err := os.OpenFile(taskFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1);
	}
	defer f.Close()

	set := new(task.Set)
	if _, err := set.ReadFrom(f); err != nil {
		fmt.Fprintf(os.Stderr, "cannot read and parse %s: %s\n", taskFileName, err)
		os.Exit(1)
	}

	if !set.Update(id, args[1]) {
		fmt.Fprintf(os.Stderr, "cannot find task with id: %v\n", id)
		os.Exit(1)
	}

	f.Seek(0, 0)
	n, err := set.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create and write json output: %s\n", err)
		os.Exit(1)
	}
	f.Truncate(n)
}

func parseId(arg string) (id uint) {
	if id64, err := strconv.ParseUint(arg, 10, 32); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse id: %s\n", err)
		os.Exit(1)
	} else {
		id = uint(id64)
	}
	return
}

func remove(args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "invalid number of arguments. usage:\n")
		fmt.Fprintf(os.Stderr, "    delete <id>\n")
		os.Exit(1)
	}

	id := parseId(args[0])

	f, err := os.OpenFile(taskFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1);
	}
	defer f.Close()

	set := new(task.Set)
	if _, err := set.ReadFrom(f); err != nil {
		fmt.Fprintf(os.Stderr, "cannot read and parse %s: %s\n", taskFileName, err)
		os.Exit(1)
	}

	if !set.Delete(id) {
		fmt.Fprintf(os.Stderr, "cannot find task with id: %v\n", id)
		os.Exit(1)
	}

	f.Seek(0, 0)
	n, err := set.WriteTo(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot create and write json output: %s\n", err)
		os.Exit(1)
	}
	f.Truncate(n)
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
	case "update":
		update(args[1:])
	case "delete":
		remove(args[1:])
	case "mark-in-progress":
		mark(task.InProgress, args[1:])
	case "mark-done":
		mark(task.Done, args[1:])
	}
}
