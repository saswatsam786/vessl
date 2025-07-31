.PHONY: build install test clean release

# Build the application
build:
	@echo "Building Vessl..."
	@./build.sh

# Install locally
install: build
	@echo "Installing Vessl..."
	@sudo cp build/vessl-$(shell go env GOOS)-$(shell go env GOARCH) /usr/local/bin/vessl
	@echo "✅ Vessl installed successfully!"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf build/
	@go clean

# Create a new release
release:
	@echo "Creating release..."
	@git tag v1.0.0
	@git push origin v1.0.0
	@echo "✅ Release v1.0.0 created!"

# Run the application
run:
	@go run main.go

# Show help
help:
	@echo "Available commands:"
	@echo "  build    - Build the application"
	@echo "  install  - Install locally"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  release  - Create a new release"
	@echo "  run      - Run the application"
	@echo "  help     - Show this help" 