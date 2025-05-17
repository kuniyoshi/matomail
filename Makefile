.PHONY: all build run fmt vet staticcheck lint clean test

# バイナリ名
BINARY_NAME=matomail
BIN_DIR=bin
CMD_DIR=cmd

all: lint build

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./$(CMD_DIR)

run: build
	$(BIN_DIR)/$(BINARY_NAME)

fmt:
	go fmt ./$(CMD_DIR)/...

vet:
	go vet ./$(CMD_DIR)/...

staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
		echo "Installing staticcheck..."; \
		go install honnef.co/go/tools/cmd/staticcheck@latest; \
	fi
	staticcheck ./$(CMD_DIR)/...

lint: fmt vet staticcheck

clean:
	go clean
	rm -rf $(BIN_DIR)

test:
	go test -v ./$(CMD_DIR)/...
