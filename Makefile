.PHONY: build test lint run clean install uninstall

BINARY=faman
CMD=./cmd/faman
PREFIX?=/usr/local
SHAREDIR=$(PREFIX)/share/faman

build:
	go build -ldflags="-s -w" -o $(BINARY) $(CMD)

test:
	go test -race ./...

lint:
	go vet ./...
	@which golangci-lint > /dev/null && golangci-lint run || echo "golangci-lint not installed"

run:
	go run $(CMD) $(ARGS)

install: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 755 $(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)
	install -d $(DESTDIR)$(SHAREDIR)
	cp -r pages $(DESTDIR)$(SHAREDIR)/
	@echo "Installed to $(PREFIX)/bin/$(BINARY)"
	@echo "Pages installed to $(SHAREDIR)/pages"

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BINARY)
	rm -rf $(DESTDIR)$(SHAREDIR)
	@echo "Uninstalled faman"

clean:
	rm -f $(BINARY) $(BINARY).exe coverage.out

pages-count:
	@ls pages/fa/*.md | wc -l
