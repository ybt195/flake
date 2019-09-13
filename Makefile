.PHONY: default clean build test coverage validate install

GO111MODULE=on

default: build

clean:
	go clean -i
	rm -f flake-coverage.out

build:
	go build ./...

test: build
	go test -v ./...

coverage: build
	go test -v -coverprofile=flake-coverage.out -covermode count -cover ./...
	go tool cover -html=flake-coverage.out
	rm flake-coverage.out

validate: test
	golangci-lint run

install: validate
	go install -i ./cmd/flake
