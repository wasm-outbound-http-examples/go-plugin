//go:build tinygo.wasm

package main

import (
	"context"

	"github.com/wasm-outbound-http-examples/go-plugin/protobufs"
)

func main() {
	protobufs.RegisterFetchService(FetchPlugin{})
}

type FetchPlugin struct{}

var _ protobufs.FetchService = (*FetchPlugin)(nil)

func (m FetchPlugin) Fetch(ctx context.Context, req *protobufs.HttpGetRequest) (*protobufs.HttpGetResponse, error) {
	hostFunctions := protobufs.NewHostFunctions()

	return hostFunctions.HttpGet(ctx, req)
}
