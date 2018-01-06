all: nyanko

nyanko: nyanko.go
	go vet -v
	go build -o nyanko -ldflags "-X main.version=$(shell git describe --tags)"

clean:
	rm -rf nyanko
