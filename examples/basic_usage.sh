#!/bin/bash

# Todo CLI - Basic Usage Examples
# 
# This script demonstrates the basic functionality of the todo CLI application.
# Make sure you have built the todo binary first: go build -o bin/todo ./cmd/todo

# Set the path to the todo binary (adjust as needed)
TODO_BIN="../bin/todo"

echo "=== Todo CLI - Basic Usage Examples ==="
echo

# Check if binary exists
if [ ! -f "$TODO_BIN" ]; then
    echo "Error: Todo binary not found at $TODO_BIN"
    echo "Please build it first with: go build -o bin/todo ./cmd/todo"
    exit 1
fi

echo "1. Adding tasks:"
$TODO_BIN add "Learn Go programming"
$TODO_BIN add "Build a CLI application"
$TODO_BIN add "Write comprehensive tests"
$TODO_BIN add "Deploy to production"
echo

echo "2. Listing all tasks:"
$TODO_BIN list
echo

echo "3. Completing some tasks:"
$TODO_BIN complete 1
$TODO_BIN done 3  # Alternative command
echo

echo "4. Editing a task:"
$TODO_BIN edit 2 "Build an awesome CLI application"
echo

echo "5. Current status:"
$TODO_BIN
echo

echo "6. Marking a task as incomplete again:"
$TODO_BIN undo 1
echo

echo "7. Final status:"
$TODO_BIN list
echo

echo "8. Using short aliases:"
$TODO_BIN a "Quick task with alias"  # 'a' for add
$TODO_BIN c 5                        # 'c' for complete
$TODO_BIN ls                         # 'ls' for list
echo

echo "9. Cleaning up - delete a task:"
$TODO_BIN delete 5
echo

echo "10. Final list:"
$TODO_BIN
echo

echo "=== Examples completed! ==="
echo "Try running '$TODO_BIN -i' for interactive mode."
echo "Use '$TODO_BIN help' to see all available commands."