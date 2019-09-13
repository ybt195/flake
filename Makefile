.PHONY: default clean build build-docker test coverage validate package

GO111MODULE=on
VERSION=$(shell cat VERSION)
LDFLAGS="-X github.com/ybt195/flake/cmd.Version=${VERSION} -X github.com/ybt195/flake/cmd.Commit=$(shell git rev-parse HEAD) -X github.com/ybt195/flake/cmd.Date=$(shell date +%Y-%m-%d.%H:%M:%S)"

default: build

clean:
	go clean -i
	rm -f flake-coverage.out
	rm -rf dist/
	rm -rf vendor/

build:
	go build ./...

build-docker:
	docker build -t ybt195/flake:$(VERSION) .

test: build
	go test -v ./...

coverage: build
	go test -v -coverprofile=flake-coverage.out -covermode count -cover ./...
	go tool cover -html=flake-coverage.out
	rm flake-coverage.out

validate: test
	golangci-lint run

package: validate
	go build -ldflags ${LDFLAGS} -o ./dist/flake ./cmd/flake
