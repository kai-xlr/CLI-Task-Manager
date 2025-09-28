package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kai-xlr/CLI-Task-Manager/internal/todo"
)

const (
	// Default filename for storing todos
	todoFile = "todos.json"
	// Version of the application
	version = "1.0.0"
)

// Configuration holds the application configuration
type Config struct {
	TodoFile    string
	Interactive bool
	Help        bool
	Version     bool
}

func main() {
	config := parseFlags()

	// Handle version flag
	if config.Version {
		fmt.Printf("todo version %s\n", version)
		return
	}

	// Show help if requested
	if config.Help {
		printHelp()
		return
	}

	// Initialize and load todo list
	todoList := todo.NewList()
	if err := loadTodosIfExists(todoList, config.TodoFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading todos: %v\n", err)
		os.Exit(1)
	}

	// Handle interactive mode
	if config.Interactive {
		runInteractive(todoList, config.TodoFile)
		return
	}

	// Handle command line arguments
	args := flag.Args()
	if len(args) == 0 {
		// Default action: print the todo list
		fmt.Print(todoList)
		return
	}

	// Execute the specified command
	if err := executeCommand(todoList, config.TodoFile, args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// parseFlags parses command line flags and returns configuration
func parseFlags() *Config {
	config := &Config{}
	
	flag.BoolVar(&config.Interactive, "i", false, "Run in interactive mode")
	flag.BoolVar(&config.Interactive, "interactive", false, "Run in interactive mode")
	flag.BoolVar(&config.Help, "h", false, "Show help message")
	flag.BoolVar(&config.Help, "help", false, "Show help message")
	flag.BoolVar(&config.Version, "v", false, "Show version information")
	flag.BoolVar(&config.Version, "version", false, "Show version information")
	flag.StringVar(&config.TodoFile, "f", todoFile, "Todo file path")
	flag.StringVar(&config.TodoFile, "file", todoFile, "Todo file path")

	flag.Parse()
	return config
}

// loadTodosIfExists loads todos from file if it exists
func loadTodosIfExists(todoList *todo.List, filename string) error {
	if _, err := os.Stat(filename); err == nil {
		return todoList.Load(filename)
	}
	return nil
}

// executeCommand executes the specified command with arguments
func executeCommand(todoList *todo.List, filename string, args []string) error {
	command := strings.ToLower(args[0])

	switch command {

	case "add", "a":
		return handleAdd(todoList, filename, args[1:])
		
	case "list", "ls", "l":
		return handleList(todoList)
		
	case "complete", "done", "c":
		return handleComplete(todoList, filename, args[1:])
		
	case "uncomplete", "undo", "u":
		return handleUncomplete(todoList, filename, args[1:])
		
	case "delete", "remove", "rm", "d":
		return handleDelete(todoList, filename, args[1:])
		
	case "edit", "e":
		return handleEdit(todoList, filename, args[1:])
		
	case "clear":
		return handleClear(todoList, filename)
		
	case "help", "h":
		printHelp()
		return nil
		
	default:
		return fmt.Errorf("unknown command: %s\nRun 'todo help' for usage information", command)
	}
}

// Command handlers

// handleAdd adds a new todo item
func handleAdd(todoList *todo.List, filename string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing todo text")
	}

	text := strings.Join(args, " ")
	index := todoList.Add(text)
	if err := saveTodos(todoList, filename); err != nil {
		return err
	}

	fmt.Printf("Added: %s (item #%d)\n", text, index+1)
	return nil
}

// handleList displays the todo list
func handleList(todoList *todo.List) error {
	fmt.Print(todoList)
	return nil
}

// handleComplete marks an item as completed
func handleComplete(todoList *todo.List, filename string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing item number")
	}

	index, err := parseItemNumber(args[0])
	if err != nil {
		return err
	}

	if err := todoList.Complete(index); err != nil {
		return err
	}

	if err := saveTodos(todoList, filename); err != nil {
		return err
	}

	fmt.Printf("Marked item #%d as completed\n", index+1)
	return nil
}

// handleUncomplete marks an item as not completed
func handleUncomplete(todoList *todo.List, filename string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing item number")
	}

	index, err := parseItemNumber(args[0])
	if err != nil {
		return err
	}

	if err := todoList.Uncomplete(index); err != nil {
		return err
	}

	if err := saveTodos(todoList, filename); err != nil {
		return err
	}

	fmt.Printf("Marked item #%d as not completed\n", index+1)
	return nil
}

// handleDelete removes an item from the list
func handleDelete(todoList *todo.List, filename string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing item number")
	}

	index, err := parseItemNumber(args[0])
	if err != nil {
		return err
	}

	// Get the item text before deleting for confirmation message
	if index >= todoList.Count() {
		return fmt.Errorf("item index out of range")
	}
	itemText := todoList.Items[index].Text

	if err := todoList.Delete(index); err != nil {
		return err
	}

	if err := saveTodos(todoList, filename); err != nil {
		return err
	}

	fmt.Printf("Deleted: %s\n", itemText)
	return nil
}

// handleEdit updates the text of an existing item
func handleEdit(todoList *todo.List, filename string, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("missing item number and/or new text")
	}

	index, err := parseItemNumber(args[0])
	if err != nil {
		return err
	}

	newText := strings.Join(args[1:], " ")
	if err := todoList.Edit(index, newText); err != nil {
		return err
	}

	if err := saveTodos(todoList, filename); err != nil {
		return err
	}

	fmt.Printf("Updated item #%d: %s\n", index+1, newText)
	return nil
}

// handleClear removes all items from the list
func handleClear(todoList *todo.List, filename string) error {
	count := todoList.Count()
	if count == 0 {
		fmt.Println("Todo list is already empty")
		return nil
	}

	todoList.Clear()
	if err := saveTodos(todoList, filename); err != nil {
		return err
	}

	fmt.Printf("Cleared %d item(s) from the todo list\n", count)
	return nil
}

// Helper functions

// parseItemNumber parses and validates an item number from string
func parseItemNumber(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid item number: %s", s)
	}

	if num < 1 {
		return 0, fmt.Errorf("item number must be greater than 0")
	}

	return num - 1, nil // Convert to 0-based index
}

// saveTodos saves the todo list to file
func saveTodos(list *todo.List, filename string) error {
	if err := list.Save(filename); err != nil {
		return fmt.Errorf("failed to save todos: %w", err)
	}
	return nil
}

func printHelp() {
	helpText := fmt.Sprintf(`Todo - A simple and efficient command line task manager

Usage:
  todo [flags] [command] [arguments]

Flags:
  -h, --help           Show this help message
  -v, --version        Show version information
  -i, --interactive    Run in interactive mode
  -f, --file <path>    Specify todo file path (default: %s)

Commands:
  add, a <text>        Add a new todo item
  list, ls, l          List all todo items (default when no command given)
  complete, done, c <n>    Mark item n as completed
  uncomplete, undo, u <n>  Mark item n as not completed
  delete, remove, rm, d <n>  Delete item n
  edit, e <n> <text>   Edit item n with new text
  clear                Clear all items
  help, h              Show this help message

Examples:
  todo add "Learn Go testing"     # Add a new task
  todo list                       # List all tasks
  todo complete 2                 # Mark task 2 as completed
  todo edit 1 "Updated task"       # Edit task 1
  todo delete 3                   # Delete task 3
  todo -i                         # Start interactive mode
  todo -f my-tasks.json list      # Use custom file

For more information, visit: https://github.com/kai-xlr/CLI-Task-Manager
`, todoFile)
	fmt.Print(helpText)
}

func runInteractive(list *todo.List, filename string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Todo Interactive Mode (v%s)\n", version)
	fmt.Println("Type 'help' for available commands or 'quit' to exit.")

	for {
		fmt.Printf("\n%s\n", list)
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		cmd := strings.ToLower(parts[0])

		switch cmd {
		case "add", "a":
			if len(parts) < 2 {
				fmt.Println("Error: missing todo text")
				continue
			}
			text := strings.Join(parts[1:], " ")
			if err := handleAdd(list, filename, []string{text}); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "list", "ls", "l":
			// List is shown at the top of each loop, so just continue
			continue

		case "complete", "done", "c":
			if len(parts) < 2 {
				fmt.Println("Error: missing item number")
				continue
			}
			if err := handleComplete(list, filename, parts[1:]); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "uncomplete", "undo", "u":
			if len(parts) < 2 {
				fmt.Println("Error: missing item number")
				continue
			}
			if err := handleUncomplete(list, filename, parts[1:]); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "delete", "remove", "rm", "d":
			if len(parts) < 2 {
				fmt.Println("Error: missing item number")
				continue
			}
			if err := handleDelete(list, filename, parts[1:]); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "edit", "e":
			if len(parts) < 3 {
				fmt.Println("Error: missing item number and/or new text")
				continue
			}
			if err := handleEdit(list, filename, parts[1:]); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "clear":
			if err := handleClear(list, filename); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "help", "h":
			printInteractiveHelp()

		case "quit", "exit", "q":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}

// printInteractiveHelp shows help specific to interactive mode
func printInteractiveHelp() {
	helpText := `Available commands in interactive mode:

  add, a <text>        Add a new todo item
  complete, done, c <n>    Mark item n as completed
  uncomplete, undo, u <n>  Mark item n as not completed
  delete, remove, rm, d <n>  Delete item n
  edit, e <n> <text>   Edit item n with new text
  clear                Clear all items
  list, ls, l          Show todo list (default view)
  help, h              Show this help message
  quit, exit, q        Exit interactive mode

Examples:
  add Buy milk
  complete 1
  edit 2 Updated task text
  delete 3
`
	fmt.Print(helpText)
}
