all: nyanko

APP		=	nyanko
VERSION	=	$(shell git describe --tags)

nyanko: nyanko.go
	go vet -v
	go build -o nyanko -ldflags "-X main.version=$(version)"

release: all
	mkdir -p $(APP)-$(VERSION)
	cp nyanko $(APP)-$(VERSION)
	zip -r $(APP)-$(VERSION).zip $(APP)-$(VERSION)
	rm -rf $(APP)-$(VERSION)

clean:
	rm -rf nyanko
