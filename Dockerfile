FROM golang:1.12.9-alpine3.10 AS builder

WORKDIR /go/src/github.com/ybt195/flake

COPY ./scripts/setup ./scripts/setup
RUN ./scripts/setup/setup.sh

ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make package

FROM alpine:3.10.2
LABEL maintainer="Jonathan Ben-tzur <jonathan.bentzur@gmail.com>"
COPY --from=builder /go/src/github.com/ybt195/flake/dist/flake /bin/flake
ENTRYPOINT ["flake"]
