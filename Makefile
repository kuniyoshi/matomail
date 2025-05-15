.PHONY: all build run fmt vet staticcheck lint clean test

# バイナリ名
BINARY_NAME=matomail

all: lint build

build:
	go build -o $(BINARY_NAME) .

run: build
	./$(BINARY_NAME)

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
	rm -f $(BINARY_NAME)

test:
	go test -v ./...
