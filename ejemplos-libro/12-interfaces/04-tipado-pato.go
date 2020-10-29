package main

import (
	"fmt"
	"time"
)

type ConversorMilisegundos interface {
	Milliseconds() int64
}

func MuestraMS(c ConversorMilisegundos) {
	fmt.Println("esto son", c.Milliseconds(), "milisegundos")
}

func main() {
	// el tipo time.Duration, que ya existe, se amolda a la
	// interfaz ConversorMilisegundos, recién creada
	MuestraMS(23 * time.Second)
}
