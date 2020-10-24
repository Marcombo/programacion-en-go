package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rest struct {
	entradas map[string]string
}

func (mr *Rest) ServeHTTP(
	rw http.ResponseWriter, req *http.Request) {

	// obtener identificador y documento (si hay alguno)
	// de la petición
	identificador := req.URL.Path
	var documento string
	if req.Body != nil {
		b, _ := ioutil.ReadAll(req.Body)
		documento = string(b)
	}

	// las respuestas a retornar no siguen un formato concreto
	rw.Header().Add("Content-Type", "text/plain")

	// según el método de la petición, nuestro http.Handler
	// llevará a cabo una u otra acción
	switch req.Method {
	case http.MethodGet:
		mr.peticionGet(identificador, rw)
	case http.MethodDelete:
		mr.peticionDelete(identificador, rw)
	case http.MethodPut:
		mr.peticionPut(identificador, documento, rw)
	case http.MethodPost:
		mr.peticionPost(identificador, documento, rw)
	default:
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "inválido:", req.Method)
	}
}

func (mr *Rest) peticionGet(
	identificador string, rw http.ResponseWriter) {

	if documento, ok := mr.entradas[identificador]; ok {
		fmt.Fprintln(rw, documento)
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "no encontrado:", identificador)
	}
}

func (mr *Rest) peticionDelete(
	identificador string, rw http.ResponseWriter) {
	if _, ok := mr.entradas[identificador]; ok {
		delete(mr.entradas, identificador)
		fmt.Fprintln(rw, "OK")
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "no encontrado:", identificador)
	}
}

func (mr *Rest) peticionPut(
	identificador, documento string, rw http.ResponseWriter) {

	if _, ok := mr.entradas[identificador]; ok {
		mr.entradas[identificador] = documento
		fmt.Fprintln(rw, "OK")
	} else {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "no encontrado:", identificador)
	}
}

func (mr *Rest) peticionPost(
	identificador, documento string, rw http.ResponseWriter) {

	if _, ok := mr.entradas[identificador]; ok {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "ya existente:", identificador)
	} else {
		mr.entradas[identificador] = documento
		fmt.Fprintln(rw, "OK")
	}
}

func main() {
	capitales := Rest{
		entradas: map[string]string{},
	}

	http.Handle("/", &capitales)

	panic(http.ListenAndServe(":8080", nil))
}
