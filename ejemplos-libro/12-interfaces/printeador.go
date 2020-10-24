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
	MuestraMS(23 * time.Second)
}
