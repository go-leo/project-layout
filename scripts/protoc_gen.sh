#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ ! $(command -v protoc-gen-go) ]
then
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc-gen-go --version
fi

if [ ! $(command -v protoc-gen-go-grpc) ]
then
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc-gen-go-grpc --version
fi

if [ ! $(command -v protoc-gen-validate-go) ]
then
	go install github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go@latest
fi

if [ ! $(command -v protoc-gen-go-leo) ]
then
	go install github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo@latest
	protoc-gen-go-leo --version
fi


PROTO_FILES=$(find . -name "*.proto")

echo "--- protoc generate start ---"
protoc \
  --proto_path=. \
  --go_out=. \
  --go_opt=paths=source_relative \
  --validate_out=. \
  --validate_opt=paths=source_relative,lang=go \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  --go-leo_out=. \
  --go-leo_opt=paths=source_relative \
  ${PROTO_FILES[*]}
echo "--- protoc generate end ---"
