all: main

main: main.go
	go vet -v
	go build -o main -ldflags "-X main.version=$(shell git describe --tags)"

clean:
	rm -rf main
