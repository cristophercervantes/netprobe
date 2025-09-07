# xprobe Makefile
# Author: Christopher

BINARY=xprobe
VERSION=1.0
AUTHOR=Christopher

.PHONY: all build install clean

all: build

build:
	@echo "Building xprobe v${VERSION}..."
	go build -o ${BINARY} .

install: build
	@echo "Installing xprobe to /usr/local/bin..."
	sudo cp ${BINARY} /usr/local/bin/

clean:
	@echo "Cleaning up..."
	rm -f ${BINARY}

test:
	@echo "Testing xprobe..."
	go test -v

version:
	@echo "xprobe v${VERSION} by ${AUTHOR}"
