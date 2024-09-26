# Labyrinth Game in Go and Ebitengine

## Prerequisites

Install Go by following the instructions at https://go.dev/doc/install .

Note that on macOS you can install Go using Homebrew.

## Set up development environment

Add the dependencies by running

```bash
go mod tidy
```

## Run the game

Run the game by running

```bash
go run .
```


## Run the Tests

```bash
go test
```

## Build for deployment

From https://ebitengine.org/en/documents/webassembly.html

```shell
env GOOS=js GOARCH=wasm go build -o ./build/smcd-go-labyrinth.wasm ./
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build/
cp ./index.html ./build/index.html
zip -r9 smcd-go-labyrinth.zip ./build/
```

## Links

- https://ebitengine.org/
- https://go.dev/doc/install
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
