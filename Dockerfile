FROM golang:1.12.9-alpine3.10 AS builder

RUN apk add --no-cache git mercurial

WORKDIR /go/src/github.com/ybt195/flake

ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o ./build/flake ./cmd/flake

FROM alpine:3.10.2
LABEL maintainer="Jonathan Ben-tzur <jonathan.bentzur@gmail.com>"
COPY --from=builder /go/src/github.com/ybt195/flake/build/flake /bin/flake
ENTRYPOINT ["flake"]
