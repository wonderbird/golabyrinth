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
env GOOS=js GOARCH=wasm go build -o ./build/raw/smcd-go-labyrinth.wasm ./
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build/raw/
cp ./index.html ./build/raw/index.html
zip -r9 ./build/smcd-go-labyrinth.zip ./build/raw
```

## Serve index.html locally

This requires either npm serve or a build in server like those offered by JetBrains tooling.

### Disclaimer

If you want to install a local http server use the following commands:

```shell
npm install --global serve
serve ./build/raw
```

Open the shown localhost link. (default port 3000 will switch to different open port if occupied)

## Links

- https://ebitengine.org/
- https://go.dev/doc/install
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
