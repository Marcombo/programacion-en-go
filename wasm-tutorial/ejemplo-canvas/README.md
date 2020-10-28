# Ejemplo Go, WASM y Canvas

Para ejecutarlo:

```go
GOOS=js GOARCH=wasm go build -o site/main.wasm src/main.go
go run server.go
```

Y abre la siguiente URL en el navegador [http://localhost:8080](http://localhost:8080).