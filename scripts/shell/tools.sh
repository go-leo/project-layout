#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

echo "--- install tools start ---"
go install github.com/google/wire/cmd/wire@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/golang/mock/mockgen@latest
go install mvdan.cc/gofumpt@latest
echo "--- install tools end ---"
