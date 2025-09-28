# Todo - CLI Task Manager

[![Go Version](https://img.shields.io/badge/go-1.24%2B-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](#license)

A simple, efficient, and user-friendly command-line task management application written in Go. Manage your todos with ease using either command-line interface or interactive mode.

## Features

### ✅ Core Functionality
- **Add tasks** with simple commands or interactive mode
- **Complete/Uncomplete tasks** to track your progress
- **Edit tasks** to update text as needed
- **Delete tasks** when they're no longer relevant
- **List tasks** with completion status and statistics
- **Clear all tasks** when starting fresh
- **Persistent storage** with automatic JSON file management

### ✅ User Experience
- **Interactive mode** for continuous task management
- **Command aliases** for faster typing (e.g., `a` for `add`, `c` for `complete`)
- **Smart error handling** with helpful error messages
- **Progress tracking** with completion statistics
- **Timestamps** for task creation and completion
- **Custom file support** with `-f` flag

### ✅ Developer Experience
- **Clean architecture** following Go best practices
- **Comprehensive test coverage** for reliability
- **Modular design** with separate packages
- **Well-documented code** with clear API

## Quick Start

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/kai-xlr/CLI-Task-Manager.git
   cd CLI-Task-Manager
   ```

2. **Build the application:**
   ```bash
   go build -o bin/todo ./cmd/todo
   ```

3. **Add to PATH (optional):**
   ```bash
   # Linux/macOS
   sudo cp bin/todo /usr/local/bin/
   
   # Or add the bin directory to your PATH
   export PATH="$PWD/bin:$PATH"
   ```

### Basic Usage

```bash
# Add a new task
todo add "Learn Go programming"

# List all tasks
todo list

# Complete a task
todo complete 1

# Start interactive mode
todo -i
```

## Usage Guide

### Command Line Mode

```bash
# Adding tasks
todo add "Buy groceries"
todo a "Call mom"                    # Short alias

# Viewing tasks
todo list                            # or just 'todo'
todo ls                              # Short alias

# Managing tasks
todo complete 1                      # Mark task 1 as completed
todo done 2                          # Alternative command
todo uncomplete 1                    # Mark as not completed
todo undo 1                          # Alternative command

# Editing tasks
todo edit 1 "Updated task text"
todo e 2 "New description"           # Short alias

# Deleting tasks
todo delete 3                        # Delete task 3
todo rm 4                            # Short alias

# Utility commands
todo clear                           # Remove all tasks
todo help                            # Show help
todo version                         # Show version
```

### Interactive Mode

Start interactive mode for continuous task management:

```bash
todo -i
```

In interactive mode, you can use all commands without the `todo` prefix:

```
Todo Interactive Mode (v1.0.0)
Type 'help' for available commands or 'quit' to exit.

Todo List (0/2 completed):
1. [ ] Learn Go programming
2. [✓] Buy groceries

> add "Study for exam"
Added: Study for exam (item #3)

> complete 1
Marked item #1 as completed

> quit
Goodbye!
```

### Custom File Location

Use a different file for your todos:

```bash
todo -f work-tasks.json add "Finish project"
todo --file personal.json list
```

## Project Architecture

The project follows Go's standard project layout with clean separation of concerns:

```
.
├── cmd/todo/           # Application entry point and CLI handling
│   └── main.go        # Main application logic and command routing
├── internal/todo/      # Internal application logic
│   ├── todo.go        # Core todo item and list functionality
│   └── todo_test.go   # Comprehensive unit tests
├── bin/               # Compiled binaries (created during build)
├── go.mod             # Go module definition
├── .gitignore         # Git ignore rules
└── README.md          # This documentation
```

### Core Components

- **`internal/todo/todo.go`**: Core functionality with `Item` and `List` types
  - `Item`: Todo item with text, completion status, and timestamps
  - `List`: Todo list with CRUD operations, statistics, and persistence
- **`internal/todo/todo_test.go`**: Comprehensive unit tests covering all functionality
- **`cmd/todo/main.go`**: CLI application with command routing, flags, and interactive mode

## Command Reference

### Available Commands

| Command | Aliases | Description | Example |
|---------|---------|-------------|----------|
| `add` | `a` | Add a new todo item | `todo add "Buy milk"` |
| `list` | `ls`, `l` | List all todo items | `todo list` |
| `complete` | `done`, `c` | Mark item as completed | `todo complete 1` |
| `uncomplete` | `undo`, `u` | Mark item as not completed | `todo undo 1` |
| `delete` | `remove`, `rm`, `d` | Delete an item | `todo delete 2` |
| `edit` | `e` | Edit item text | `todo edit 1 "New text"` |
| `clear` | | Remove all items | `todo clear` |
| `help` | `h` | Show help message | `todo help` |
| `version` | `v` | Show version info | `todo version` |

### Available Flags

| Flag | Long Form | Description | Example |
|------|-----------|-------------|----------|
| `-h` | `--help` | Show help message | `todo -h` |
| `-v` | `--version` | Show version info | `todo -v` |
| `-i` | `--interactive` | Start interactive mode | `todo -i` |
| `-f` | `--file` | Specify todo file path | `todo -f tasks.json list` |

## Prerequisites

- **Go 1.24.5 or later** for building from source
- **Git** for cloning the repository

## Data Storage

Todos are automatically saved to `todos.json` in the current directory. The JSON format includes:

```json
{
  "items": [
    {
      "text": "Learn Go programming",
      "done": false,
      "created_at": "2024-01-15T10:30:00Z"
    },
    {
      "text": "Buy groceries",
      "done": true,
      "created_at": "2024-01-15T09:15:00Z",
      "completed_at": "2024-01-15T11:45:00Z"
    }
  ]
}
```

## API Usage

You can also use the todo package directly in your Go applications:

```go
package main

import (
    "fmt"
    "log"
    "github.com/kai-xlr/CLI-Task-Manager/internal/todo"
)

func main() {
    // Create a new todo list
    list := todo.NewList()
    
    // Add some items
    list.Add("Learn Go")
    list.Add("Build CLI app")
    list.Add("Write tests")
    
    // Complete an item
    if err := list.Complete(0); err != nil {
        log.Fatal(err)
    }
    
    // Display statistics
    fmt.Printf("Total: %d, Completed: %d, Pending: %d\n", 
        list.Count(), list.CountCompleted(), list.CountPending())
    
    // Save to file
    if err := list.Save("my-todos.json"); err != nil {
        log.Fatal(err)
    }
    
    // Load from file
    newList := todo.NewList()
    if err := newList.Load("my-todos.json"); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(newList)
}
```

## Development

### Getting Started

1. **Fork and clone the repository**
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Run tests to ensure everything works:**
   ```bash
   go test ./...
   ```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./internal/todo

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Code Quality

```bash
# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run linter (requires golangci-lint)
golangci-lint run

# Run all quality checks
make quality  # if Makefile is available
```

### Building

```bash
# Build for current platform
go build -o bin/todo ./cmd/todo

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o bin/todo-linux ./cmd/todo
GOOS=windows GOARCH=amd64 go build -o bin/todo.exe ./cmd/todo
GOOS=darwin GOARCH=amd64 go build -o bin/todo-macos ./cmd/todo

# Build with version info
go build -ldflags "-X main.version=v1.0.0" -o bin/todo ./cmd/todo

# Clean build artifacts
rm -rf bin/
```

### Testing the CLI

```bash
# Build and test basic functionality
go build -o bin/todo ./cmd/todo
./bin/todo add "Test task"
./bin/todo list
./bin/todo complete 1
./bin/todo -i  # Test interactive mode
```

## Contributing

We welcome contributions! This project follows standard Go practices and includes comprehensive tests.

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch:** `git checkout -b feature/amazing-feature`
3. **Make your changes and add tests**
4. **Run tests:** `go test ./...`
5. **Commit your changes:** `git commit -m 'Add amazing feature'`
6. **Push to the branch:** `git push origin feature/amazing-feature`
7. **Open a Pull Request**

### Development Guidelines

- **Write tests** for new functionality
- **Follow Go conventions** and use `go fmt`
- **Update documentation** for user-facing changes
- **Keep commits atomic** and write clear commit messages
- **Run `go vet` and `golangci-lint`** before submitting

### Potential Enhancements

Contributions are welcome for these features:

- 🚀 **Priority levels** for tasks
- 📅 **Due dates** and reminders
- 🏷️ **Tags and categories** for organization
- 🔍 **Search and filter** capabilities
- ⚙️ **Configuration file** support
- 🎨 **Colored output** and themes
- 📊 **Statistics and reporting** features
- 🔄 **Import/export** functionality
- 📱 **Mobile-friendly** TUI interface

## Troubleshooting

### Common Issues

**Build fails with Go version error:**
```bash
# Check your Go version
go version
# Upgrade Go if needed (requires Go 1.21+)
```

**Permission denied when installing:**
```bash
# Use sudo for system-wide installation
sudo cp bin/todo /usr/local/bin/
# Or install to user directory
cp bin/todo ~/.local/bin/
```

**Todo file not found:**
```bash
# Check current directory
pwd
# The todos.json file is created in the current directory
# Use -f flag to specify a different location
todo -f ~/todos.json list
```

**Interactive mode not working:**
- Ensure your terminal supports interactive input
- Try using a different terminal or shell
- Check that stdin is not being redirected

## Examples

### Daily Workflow

```bash
# Morning: Check your tasks
todo

# Add new tasks as they come up
todo add "Review pull requests"
todo add "Update documentation"
todo add "Plan team meeting"

# During the day: Complete tasks
todo complete 1
todo done 3

# Need to edit a task?
todo edit 2 "Update API documentation with examples"

# End of day: Review progress
todo list
```

### Project Management

```bash
# Use different files for different projects
todo -f work.json add "Deploy to production"
todo -f personal.json add "Book vacation"
todo -f side-project.json add "Implement user auth"

# Quick check across projects
todo -f work.json
todo -f personal.json
```

### Interactive Session Example

```
$ todo -i
Todo Interactive Mode (v1.0.0)
Type 'help' for available commands or 'quit' to exit.

Todo List (2/5 completed):
1. [✓] Set up development environment
2. [ ] Write unit tests
3. [✓] Implement core features
4. [ ] Add documentation
5. [ ] Deploy to production

> add "Code review with team"
Added: Code review with team (item #6)

> complete 2
Marked item #2 as completed

> edit 4 "Add comprehensive documentation and examples"
Updated item #4: Add comprehensive documentation and examples

> quit
Goodbye!
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Go](https://golang.org/) for performance and simplicity
- Inspired by simple, effective task management tools
- Thanks to the Go community for excellent tooling and practices

## Contact

- **Author:** Kai XLR
- **Repository:** [https://github.com/kai-xlr/CLI-Task-Manager](https://github.com/kai-xlr/CLI-Task-Manager)
- **Issues:** [Report bugs or request features](https://github.com/kai-xlr/CLI-Task-Manager/issues)

---

**Made with ❤️ and Go**
