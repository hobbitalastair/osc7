.POSIX:

PREFIX = /usr/local
BINDIR = $(PREFIX)/bin

all: osc7

osc7:
	go build -trimpath -o osc7 .

test:
	go test ./...

clean:
	rm -f osc7

install: osc7
	install -d "$(DESTDIR)$(BINDIR)"
	install -m 0755 osc7 "$(DESTDIR)$(BINDIR)/osc7"

.PHONY: all test clean install
