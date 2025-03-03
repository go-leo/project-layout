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

if [ ! $(command -v protoc-gen-go-gors) ]
then
	go install github.com/go-leo/gors/cmd/protoc-gen-go-gors@latest
	protoc-gen-go-gors --version
fi

if [ ! $(command -v protoc-gen-go-cqrs) ]
then
	go install github.com/go-leo/design-pattern/cqrs/cmd/protoc-gen-go-cqrs@latest
	protoc-gen-go-cqrs --version
fi

proto_files=$(find api -name "*.proto")

echo "--- protoc generate start ---"
protoc \
  --proto_path=. \
  --go_out=. \
  --go_opt=module=github.com/go-leo/ddd-layout \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/go-leo/ddd-layout \
  --go-grpc_opt=require_unimplemented_servers=false \
  --go-gors_out=. \
  --go-gors_opt=module=github.com/go-leo/ddd-layout \
  --go-gors_opt=path_to_lower=true \
  ${proto_files[*]}
echo "--- protoc generate end ---"
