# CLI Task Manager

A command-line task management application written in Go, designed to help you organize and track your todos efficiently.

## Project Status

🚧 **In Development** - This project is currently in early development. Basic todo item functionality has been implemented, but the CLI interface is not yet complete.

## Features

- ✅ Todo item creation with text and completion status
- ✅ Todo list management (add, complete items)
- ✅ Persistent storage (save/load to JSON files)
- ✅ Formatted string output for display
- ✅ Clean, testable architecture following Go best practices
- ✅ Comprehensive test coverage
- 🚧 Command-line interface (planned)
- 🚧 Task categorization (planned)

## Architecture

The project follows Go's standard project layout:

```
.
├── cmd/todo/           # Application entry point
├── internal/todo/      # Internal application logic
├── go.mod             # Go module definition
└── README.md          # This file
```

### Core Components

- **`internal/todo/todo.go`**: Defines the `Item` and `List` structs with full functionality
  - `Item`: Todo item with text and completion status
  - `List`: Todo list with operations for adding, completing, saving, and loading
- **`internal/todo/todo_test.go`**: Comprehensive unit tests covering all functionality
- **`cmd/todo/main.go`**: Application entry point (currently minimal, needs package main)

## Prerequisites

- Go 1.24.5 or later
- Git (for cloning the repository)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kai-xlr/CLI-Task-Manager.git
   cd CLI-Task-Manager
   ```

2. Download dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o bin/todo ./cmd/todo
   ```

   *Note: The application is not yet functional as the main.go file contains only a package declaration.*

## Usage

*The CLI interface is currently under development.* However, you can use the full todo list functionality through the Go package:

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
    
    // Display the list
    fmt.Println(list.String())
    
    // Save to file
    if err := list.Save("my-todos.json"); err != nil {
        log.Fatal(err)
    }
    
    // Load from file
    newList := todo.NewList()
    if err := newList.Load("my-todos.json"); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Loaded list:")
    fmt.Println(newList.String())
}
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./internal/todo

# Run specific test functions
go test -run TestNewItem ./internal/todo
go test -run TestAddItem ./internal/todo
go test -run TestCompleteItem ./internal/todo
go test -run TestSaveAndLoad ./internal/todo
```

### Code Quality

```bash
# Format code
go fmt ./...

# Vet code for potential issues
go vet ./...

# Run linter (requires golangci-lint)
golangci-lint run
```

### Building

```bash
# Build for current platform
go build -o bin/todo ./cmd/todo

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o bin/todo-linux ./cmd/todo
GOOS=windows GOARCH=amd64 go build -o bin/todo.exe ./cmd/todo

# Clean build artifacts
rm -rf bin/
```

## Contributing

This project is in active development. The core functionality is implemented and well-tested. Key areas that need implementation:

1. **CLI Interface**: Implement the main.go file with command-line argument parsing (currently has wrong package declaration)
2. **Additional Operations**: Implement delete, edit, and filter operations
3. **Configuration**: Add configuration file support
4. **Enhanced Features**: Priority levels, due dates, categories
5. **Documentation**: Continue expanding documentation as features are added

### Current Status
- ✅ Core data structures (Item, List)
- ✅ Basic operations (Add, Complete)
- ✅ Persistence (Save/Load JSON)
- ✅ Comprehensive test coverage
- 🚧 CLI interface

## Project Structure

- **`cmd/`**: Application entry points
- **`internal/`**: Private application and library code
- **`bin/`**: Compiled binaries (not tracked in git)

## License

*License information to be added*

## Contact

*Contact information to be added*
