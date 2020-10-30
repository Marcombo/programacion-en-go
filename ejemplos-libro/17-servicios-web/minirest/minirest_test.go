package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Testing: ver capítulo 20
func TestPost(t *testing.T) {
	// DADO un servicio REST de guardado de documentos
	handler := Rest{entradas: map[string]string{}}
	servidor := httptest.NewServer(&handler)
	cliente := servidor.Client()

	// CUANDO se inserta un documento
	resp, err := cliente.Post(
		servidor.URL+"/Japon",
		"text/plain",
		strings.NewReader("Tokyo"),
	)
	if err != nil {
		t.Error(err)
	}
	// ENTONCES el servidor responde 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Error("Esperaba 200 OK pero obtuve", resp.StatusCode)
	}

	// Y CUANDO se busca de nuevo ese documento
	resp, err = cliente.Get(servidor.URL + "/Japon")
	if err != nil {
		t.Error(err)
	}
	documento, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	// ENTONCES el documento retornado es igual
	// al documento enviado (quitando el salto de línea
	// "\n" que el servidor añade)
	documento = bytes.Trim(documento, "\n")
	if string(documento) != "Tokyo" {
		t.Error("Esperaba Tokyo pero obtuve", string(documento))
	}
}
