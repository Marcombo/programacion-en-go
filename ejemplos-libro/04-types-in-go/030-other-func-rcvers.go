package main

import (
	"fmt"
)

type PH float32

func (p PH) Categoria() string {
	switch {
	case p < 7:
		return "ácido"
	case p > 7:
		return "básico"
	default:
		return "neutro"
	}
}

func main() {
	phs := []PH{PH(7), PH(1.2), PH(9)}

	for _, ph := range phs {
		fmt.Printf("Un pH == %v es %v\n", ph, ph.Categoria())
	}
}
