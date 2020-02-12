.PHONY: build clean test lint dep-update

VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
            echo v0)

build:
	export GOFLAGS=-mod=vendor
	CGO_ENABLED=0 gox -mod vendor -ldflags '-extldflags "-static" -X "main.version=${VERSION}"' .

clean:
	rm fortlit-go_*

test:
	export GOFLAGS=-mod=vendor
	go test ./...

lint:
	golangci-lint run --enable-all

dep-update:
	go get -u ./...
	go test ./...
	go mod tidy
	go mod vendor
