# Go compiler and application name
GO := go
APP := brdns

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	$(GO) build -o $(APP) main.go

# Run the application
.PHONY: run
run:
	./$(APP)

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(APP)

# Format the code
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Run linting
.PHONY: lint
lint:
	golangci-lint run

# Install dependencies
.PHONY: deps
deps:
	$(GO) mod tidy
