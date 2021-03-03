.PHONY: generate
generate:
	go generate -v

.PHONY: build
build:
	goreleaser build --rm-dist --snapshot

.PHONY: release
release:
	goreleaser release --rm-dist --snapshot --skip-publish

.PHONY: test
test:
	GOFLAGS=-mod=vendor go test -race -cover -v ./...

.PHONY: lint
lint:
	golangci-lint run --enable-all --disable gomnd --disable godox --disable exhaustivestruct --timeout 5m

.PHONY: clean
clean:
	rm -rf dist/
