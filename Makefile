all: nyanko

APP		=	nyanko
VERSION	=	$(shell git describe --tags)

nyanko: nyanko.go words.go
	go vet -v
	go build -o nyanko -ldflags "-X main.version=$(VERSION)"

words.go: words.txt
	go-bindata -o words.go words.txt

release: all
	mkdir -p $(APP)-$(VERSION)
	cp nyanko words.txt $(APP)-$(VERSION)
	zip -r $(APP)-$(VERSION).zip $(APP)-$(VERSION)
	rm -rf $(APP)-$(VERSION)

clean:
	rm -rf nyanko
