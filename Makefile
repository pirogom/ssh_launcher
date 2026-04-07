# Makefile for ssh-launcher

# Binary name
BINARY_NAME=sshl

# Build targets
.PHONY: all build build-win build-mac build-linux clean

all: build

# Build for current OS
build:
	go build -o $(BINARY_NAME) .

# Build for Windows
build-win:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe .

# Build for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-mac .

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux .

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe $(BINARY_NAME)-mac $(BINARY_NAME)-linux
