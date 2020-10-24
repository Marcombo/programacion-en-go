package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HolaServicio struct{}

func (hs *HolaServicio) ServeHTTP(
	rw http.ResponseWriter, req *http.Request) {

	rw.Header().Add("Content-Type", "text/html")

	documento := fmt.Sprintf(`
		<h1>Bienvenido!</h1>
		<p>Ruta de acceso: %s</p>
	`, req.URL.Path)
	// status code 200 es implícito si no se indica otro
	rw.Write([]byte(documento))
}

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		req, err := http.NewRequest(http.MethodGet,
			"http://localhost:8080/ruta/de/documento", nil)
		if err != nil {
			panic(err)
		}
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println("Código de respuesta:", resp.StatusCode)
		fmt.Println("Content-Type:", resp.Header["Content-Type"])
		cuerpo, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("---")
		fmt.Println(string(cuerpo))
	}()
	panic(http.ListenAndServe(":8080", &HolaServicio{}))
}
