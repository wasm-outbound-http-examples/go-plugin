package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/wasm-outbound-http-examples/go-plugin/protobufs"
)

func main() {
	ctx := context.Background()
	p, err := protobufs.NewFetchServicePlugin(ctx)
	if err != nil {
		panic(err)
	}

	fetchPlugin, err := p.Load(ctx, "plugin/plugin.wasm", myHostFunctions{})
	if err != nil {
		panic(err)
	}
	defer fetchPlugin.Close(ctx)

	response, err := fetchPlugin.Fetch(ctx, &protobufs.HttpGetRequest{
		Url: "https://httpbin.org/anything",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("text:\n" + string(response.GetResponse()))
}

type myHostFunctions struct{}

var _ protobufs.HostFunctions = (*myHostFunctions)(nil)

func (myHostFunctions) HttpGet(_ context.Context, req *protobufs.HttpGetRequest) (*protobufs.HttpGetResponse, error) {
	res, err := http.Get(req.Url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &protobufs.HttpGetResponse{Response: buf}, nil
}
