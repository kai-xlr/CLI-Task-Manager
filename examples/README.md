# Todo CLI - Examples

This directory contains examples demonstrating how to use the Todo CLI application and its underlying package.

## Examples

### 1. Basic CLI Usage (`basic_usage.sh`)

A shell script that demonstrates all the basic CLI functionality:

```bash
# Make the script executable and run it
chmod +x basic_usage.sh
./basic_usage.sh
```

This script shows:
- Adding tasks with `todo add`
- Listing tasks with `todo list` and `todo`
- Completing tasks with `todo complete` and `todo done`
- Editing tasks with `todo edit`
- Using command aliases (`a`, `c`, `ls`, etc.)
- Deleting tasks with `todo delete`

### 2. API Usage (`api_usage.go`)

A Go program that shows how to use the todo package as a library in your own applications:

```bash
# Run the API usage example
go run api_usage.go
```

This example demonstrates:
- Creating and managing todo lists programmatically
- Using all the core functionality (add, complete, edit, delete)
- Saving and loading todo lists from JSON files
- Accessing statistics and metadata

## Prerequisites

Before running the examples:

1. **Build the todo binary:**
   ```bash
   cd .. # Go back to project root
   go build -o bin/todo ./cmd/todo
   ```

2. **For the shell script example:**
   - Make sure you're on a Unix-like system (Linux, macOS, WSL)
   - Ensure the todo binary is built and accessible

3. **For the Go API example:**
   - Make sure you're in the project directory so Go can find the internal packages

## What You'll Learn

### CLI Usage
- How to perform all basic todo operations
- Command aliases for faster workflow
- Interactive vs. command-line modes
- File management with custom todo files

### Library Usage
- How to integrate the todo package into your own Go projects
- Complete API coverage with error handling
- JSON persistence and data management
- Statistics and reporting capabilities

## Next Steps

After trying these examples:

1. **Explore Interactive Mode:**
   ```bash
   ../bin/todo -i
   ```

2. **Try Custom File Locations:**
   ```bash
   ../bin/todo -f my-tasks.json add "Custom file task"
   ../bin/todo -f my-tasks.json list
   ```

3. **Build Your Own Integration:**
   Use the `api_usage.go` example as a starting point for your own applications.

## Troubleshooting

**Binary not found:**
- Make sure you've built the project: `go build -o bin/todo ./cmd/todo`
- Check the path in the shell script matches your binary location

**Permission denied (shell script):**
- Make the script executable: `chmod +x basic_usage.sh`

**Package import errors (Go example):**
- Make sure you're running from the project directory
- Ensure all dependencies are available: `go mod tidy`