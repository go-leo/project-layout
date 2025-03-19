//go:build tools
// +build tools

// following https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

package tools

import (
	_ "github.com/go-kit/kit/endpoint"
	_ "github.com/go-kit/log"
	_ "github.com/go-leo/gox/backoff"
	_ "github.com/go-leo/leo/v3"
	_ "github.com/google/wire"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/gorilla/mux"
	_ "github.com/llorllale/go-gitlint"
	_ "github.com/spf13/cobra"
	_ "github.com/ugorji/go/codec"
	_ "google.golang.org/genproto/googleapis/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	_ "google.golang.org/grpc"
	_ "google.golang.org/protobuf"
)
