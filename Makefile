.PHONY: build clean run test install

# Build settings
BINARY_NAME=gophpserver
BUILD_DIR=build
VERSION=1.0.0
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

# Default target
all: clean build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/gophpserver

# Cross-compile for multiple platforms
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/gophpserver
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/gophpserver
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/gophpserver

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

# Run the application
run:
	@go run ./cmd/gophpserver

# Install the application
install:
	@echo "Installing $(BINARY_NAME)..."
	@go install $(LDFLAGS) ./cmd/gophpserver

# Create distribution packages
dist: build-all
	@echo "Creating distribution packages..."
	@mkdir -p $(BUILD_DIR)/dist
	@cp -r public $(BUILD_DIR)/dist/
	@cp README.md $(BUILD_DIR)/dist/
	@cd $(BUILD_DIR) && tar -czf dist/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz $(BINARY_NAME)-linux-amd64 dist/public dist/README.md
	@cd $(BUILD_DIR) && tar -czf dist/$(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz $(BINARY_NAME)-darwin-amd64 dist/public dist/README.md
	@cd $(BUILD_DIR) && zip -r dist/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BINARY_NAME)-windows-amd64.exe dist/public dist/README.md