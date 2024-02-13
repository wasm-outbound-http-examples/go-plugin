# Use knqyf263/go-plugin to send HTTP(s) requests from inside WASM

This devcontainer is configured to provide you a [knqyf263/go-plugin](https://github.com/knqyf263/go-plugin)'s CLI tool installation,
Go and TinyGo toolchains, and a `protobuf-compiler`.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/wasm-outbound-http-examples/go-plugin)

The sample code and following instructions are partially based on [go-plugin's examples](https://github.com/knqyf263/go-plugin/tree/v0.8.0/examples/)
and [go-plugin's tutorial](https://github.com/knqyf263/go-plugin/blob/v0.8.0/README.md#tutorial).

## Instructions for this devcontainer

Tested with `protoc-gen-go-plugin` [0.8.0](https://github.com/knqyf263/go-plugin/releases/tag/v0.8.0), 
`protobuf-compiler` 3.12.4, Go 1.21.7, TinyGo [0.30.0](https://github.com/tinygo-org/tinygo/releases/tag/v0.30.0).

### Preparation

1. Open this repo in devcontainer, e.g. using Github Codespaces.
   Type or copy/paste following commands to devcontainer's terminal.

### Building

1. Generate the SDK codes based on interfaces declared in `protobufs/myschema.proto` file:

```sh
cd protobufs
protoc --go-plugin_out=. --go-plugin_opt=paths=source_relative myschema.proto
```

This will generate 4 new .go files in `protobufs` directory.

2. Compile the plugin sample:

```sh
cd ../plugin
tinygo build -o plugin.wasm -scheduler=none -target=wasi --no-debug plugin.go
```

This will generate `plugin.wasm` in `plugin` directory.

3. Run the main ("host") program, which makes use of the `plugin.wasm`:

```sh
cd ..
go run main.go
```

### Finish

Perform your own experiments if desired.

---

<sub>Created for (wannabe-awesome) [list](https://github.com/vasilev/HTTP-request-from-inside-WASM)</sub>
