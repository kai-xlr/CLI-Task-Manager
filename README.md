# CLI Task Manager

A command-line task management application written in Go, designed to help you organize and track your todos efficiently.

## Project Status

🚧 **In Development** - This project is currently in early development. Basic todo item functionality has been implemented, but the CLI interface is not yet complete.

## Features

- ✅ Todo item creation with text and completion status
- ✅ Clean, testable architecture following Go best practices
- 🚧 Command-line interface (planned)
- 🚧 Persistent storage (planned)
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

- **`internal/todo/todo.go`**: Defines the `Item` struct and constructor function
- **`internal/todo/todo_test.go`**: Unit tests for todo item functionality
- **`cmd/todo/main.go`**: Application entry point (currently minimal)

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

*Coming soon - The CLI interface is currently under development.*

For now, you can explore the todo item functionality through the Go package:

```go
package main

import (
    "fmt"
    "github.com/kai-xlr/CLI-Task-Manager/internal/todo"
)

func main() {
    item := todo.NewItem("Learn Go")
    fmt.Printf("Task: %s, Done: %t\n", item.Text, item.Done)
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

# Run a specific test function
go test -run TestNewItem ./internal/todo
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

This project is in early development. Key areas that need implementation:

1. **CLI Interface**: Implement the main.go file with command-line argument parsing
2. **Persistent Storage**: Add functionality to save and load todos from disk
3. **Additional Operations**: Implement list, complete, delete, and update operations
4. **Configuration**: Add configuration file support
5. **Documentation**: Expand documentation as features are added

## Project Structure

- **`cmd/`**: Application entry points
- **`internal/`**: Private application and library code
- **`bin/`**: Compiled binaries (not tracked in git)

## License

*License information to be added*

## Contact

*Contact information to be added*
