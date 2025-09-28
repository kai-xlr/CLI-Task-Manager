# Todo CLI Task Manager - Makefile

# Variables
BINARY_NAME=todo
MAIN_PACKAGE=./cmd/todo
BUILD_DIR=bin
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Todo CLI Task Manager - Available commands:"
	@echo
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Build commands
.PHONY: build
build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Built: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: install
install: build ## Install the binary to /usr/local/bin
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installed successfully!"

.PHONY: build-all
build-all: ## Build for multiple platforms
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PACKAGE)
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PACKAGE)
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PACKAGE)
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PACKAGE)
	@echo "Built binaries for multiple platforms in $(BUILD_DIR)/"

# Development commands
.PHONY: run
run: ## Run the application
	@go run $(MAIN_PACKAGE) $(ARGS)

.PHONY: dev
dev: build ## Build and run the application
	@./$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

# Testing commands
.PHONY: test
test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: benchmark
benchmark: ## Run benchmarks
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Code quality commands
.PHONY: fmt
fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

.PHONY: vet
vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

.PHONY: lint
lint: ## Run golangci-lint (requires golangci-lint to be installed)
	@echo "Running linter..."
	@golangci-lint run

.PHONY: quality
quality: fmt vet lint test ## Run all quality checks

# Utility commands
.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Cleaned!"

.PHONY: deps
deps: ## Download and tidy dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies updated!"

.PHONY: demo
demo: build ## Run a demo of the todo application
	@echo "Running todo demo..."
	@./$(BUILD_DIR)/$(BINARY_NAME) add "Learn Go programming"
	@./$(BUILD_DIR)/$(BINARY_NAME) add "Build CLI application"
	@./$(BUILD_DIR)/$(BINARY_NAME) add "Write comprehensive tests"
	@./$(BUILD_DIR)/$(BINARY_NAME) complete 1
	@./$(BUILD_DIR)/$(BINARY_NAME) list
	@echo
	@echo "Demo completed! Try running './$(BUILD_DIR)/$(BINARY_NAME) -h' for more options."

# Release commands
.PHONY: release-check
release-check: quality test ## Run pre-release checks
	@echo "Pre-release checks passed!"

# Docker commands (if Docker support is added later)
.PHONY: docker-build
docker-build: ## Build Docker image (if Dockerfile exists)
	@if [ -f Dockerfile ]; then \
		docker build -t $(BINARY_NAME):$(VERSION) .; \
	else \
		echo "Dockerfile not found"; \
	fi

.DEFAULT_GOAL := help