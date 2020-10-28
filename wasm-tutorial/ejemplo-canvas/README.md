# Go and WASM tutorial

Example code for the blog post:
[Go and WebAssembly (I): interacting with your browser JS API](http://macias.info/entry/202003151900_go_wasm_js.md).

How to run it:

```go
GOOS=js GOARCH=wasm go build -o site/main.wasm wasm-main/main.go
go run server.go
```

And open the `http://localhost:8080` URL in your browser.