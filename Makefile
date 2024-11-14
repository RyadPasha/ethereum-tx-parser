# Makefile for Go project

# Default target when `make` is run without arguments
.PHONY: run
run:
	go run main.go

# Command to build the project
.PHONY: build
build:
	go build -o ethereum-tx-parser .

# Command to run the project with a custom port
.PHONY: run-custom
run-custom:
	PORT=$(PORT) go run main.go

# Command to run tests
.PHONY: test
test:
	go test -v ./server
