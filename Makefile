.PHONY: default clean build build-docker test coverage validate package

GO111MODULE=on

default: build

clean:
	go clean -i
	rm -f flake-coverage.out
	rm -rf dist/
	rm -rf vendor/

build:
	go build ./...

build-docker:
	docker build -t ybt195/flake .

test: build
	go test -v ./...

coverage: build
	go test -v -coverprofile=flake-coverage.out -covermode count -cover ./...
	go tool cover -html=flake-coverage.out
	rm flake-coverage.out

validate: test
	golangci-lint run

package: validate
	go build -o ./dist/flake ./cmd/flake
