# Project Metadata
APP_NAME := petrock
PKG := ./...
BIN_DIR := bin
BUILD_DIR := $(BIN_DIR)/$(APP_NAME)

# Default target
.PHONY: all
all: build

# Build the app
.PHONY: build
build:
	go build -o $(BUILD_DIR) .

# Run the app
.PHONY: run
run: build
	$(BUILD_DIR)

# Test the app
.PHONY: test
test:
	go test -v $(PKG)

# Format code
.PHONY: fmt
fmt:
	go fmt $(PKG)

# Lint (requires golangci-lint)
.PHONY: lint
lint:
	golangci-lint run

# Clean built files
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# Install dependencies (for Go modules)
.PHONY: deps
deps:
	go mod tidy

# Migrate DB using golang-migrate (example)
.PHONY: migrate-up
migrate-up:
	migrate -path db/migrations -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down:
	migrate -path db/migrations -database "$(DATABASE_URL)" down

.PHONY: seed
seed:
	go run scripts/seed.go
