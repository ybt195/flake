#!/bin/sh

apk add --no-cache curl g++ git make mercurial
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin ${GOLANGCI_VERSION}
