.PHONY: all build run fmt vet staticcheck lint clean test

# バイナリ名
BINARY_NAME=matomail
BIN_DIR=bin

all: lint build

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) .

run: build
	$(BIN_DIR)/$(BINARY_NAME)

fmt:
	go fmt ./...

vet:
	go vet ./...

staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
		echo "Installing staticcheck..."; \
		go install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi
	staticcheck ./...

lint: fmt vet staticcheck

clean:
	go clean
	rm -rf $(BIN_DIR)

test:
	go test -v ./...
