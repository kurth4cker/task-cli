package main

import (
	"fmt"
	"os"
	"slices"
)

func add(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: add <task description>")
		os.Exit(1)
	}

	file, err := os.OpenFile("database.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open database: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, args[0])
	fmt.Printf("0: %s\n", args[0])
}

func list(args []string) {
	if len(args) < 1 {
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
	for _, arg := range args {
		statuses := []string{"done", "todo", "in-progress"}
		if !slices.Contains(statuses, arg) {
			os.Exit(1)
		}
	}
}
