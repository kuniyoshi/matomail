.PHONY: all build run fmt vet staticcheck lint clean test

# バイナリ名
BINARY_NAME=matomail
BIN_DIR=bin
SRC_DIR=src

all: lint build

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./$(SRC_DIR)

run: build
	$(BIN_DIR)/$(BINARY_NAME)

fmt:
	go fmt ./$(SRC_DIR)/...

vet:
	go vet ./$(SRC_DIR)/...

staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
		echo "Installing staticcheck..."; \
		go install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi
	staticcheck ./$(SRC_DIR)/...

lint: fmt vet staticcheck

clean:
	go clean
	rm -rf $(BIN_DIR)

test:
	go test -v ./$(SRC_DIR)/...
