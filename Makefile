.PHONY: default deps clean build test coverage

default: build

deps:
	dep ensure

clean:
	go clean -i
	rm -f flake-coverage.out

build: deps
	go build ./...

test: build
	go test -v ./...

coverage: build
	go test -v -coverprofile=flake-coverage.out -covermode count -cover ./...
	go tool cover -html=flake-coverage.out
	rm flake-coverage.out

install: build
	go install -i ./cmd/flake