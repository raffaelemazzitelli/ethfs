# Name of the binary
BINARY=evm
BUILD_DIR=build

.PHONY: all build run test clean init

all: build

init:
	@echo "Initializing the application..."
	@go mod init evm

build: $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)
	@echo "Building the application..."
	@go mod tidy
	@go build -o $(BUILD_DIR)/$(BINARY) main.go

run: build
	@echo "Running the application..."
	@./$(BUILD_DIR)/$(BINARY)

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

	
