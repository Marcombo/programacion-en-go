package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const respuestaCorrecta = "go build"

func main() {

	fmt.Println("Trivial Go")
	fmt.Println("==========")

	fmt.Print("Pregunta 1: qué comando se usa para generar un ejecutable? ")

	// paquete bufio: se explicará en capítulo 14
	teclado := bufio.NewReader(os.Stdin)
	respuesta, _ := teclado.ReadString('\n')
	respuesta = strings.Trim(respuesta, " \n")

	// puedes comparar cadenas con ==
	if respuesta == respuestaCorrecta {
		fmt.Println("Ganaste!")
	} else {
		fmt.Println("Perdiste :(")
	}
}
