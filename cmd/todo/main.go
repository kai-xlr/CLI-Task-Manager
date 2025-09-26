package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kai-xlr/CLI-Task-Manager/internal/todo"
)

const (
	todoFile = "todos.json"
)

func main() {
	// Load existing todos
	todoList := todo.NewList()
	if _, err := os.Stat(todoFile); err == nil {
		if err := todoList.Load(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error loading todos:", err)
			os.Exit(1)
		}
	}

	// Process command line arguments
	args := os.Args[1:] // Skip the program name

	if len(args) == 0 {
		// Default action: print the todo list
		fmt.Println(todoList)
		return
	}

	// Get the command
	command := args[0]

	if command == "add" {
		if len(args) < 2 {
			fmt.Println("Error: missing todo text")
			os.Exit(1)
		}

		// Join all remaining arguments as the todo text
		text := strings.Join(args[1:], " ")
		todoList.Add(text)
		saveTodos(todoList)
		fmt.Println("Added:", text)
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: add")
		os.Exit(1)
	}
}

func saveTodos(list *todo.List) {
	if err := list.Save(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving todos:", err)
		os.Exit(1)
	}
}
