run:
	go run ./cmd/app

build:
	mkdir -p bin
	go build -o bin ./cmd/app
