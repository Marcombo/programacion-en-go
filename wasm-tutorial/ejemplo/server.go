package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// por defecto, cargamos el archivo `index.html`
		if req.RequestURI == "/" {
			req.RequestURI = "/index.html"
		}
		// si no, el archivo que nos llegue por la URL lo iremos a buscar
		// a la carpeta "site"
		file, err := os.Open(filepath.Join("./site", req.RequestURI))
		if err == nil {
			// mandamos el contenido del archivo leído hacia la respuesta HTTP
			io.Copy(w, file)
		}
		// por sencillez, ignoramos cualquier gestión de errores HTTP
		// por ejemplo, si no se encuentra un archivo
	})
	// el servidor escuchará en el puerto 8080 de la máquina local
	fmt.Println(http.ListenAndServe(":8080", nil))
}
