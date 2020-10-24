package main

import (
	"fmt"
)

type Perro struct{}

func (p *Perro) Ladra() string {
	return "¡Guau!"
}

type Rana struct{}

func (r *Rana) Croa() string {
	return "¡Croac!"
}

type Canguro struct{}

func (c *Canguro) Salta() string {
	return "¡Boing!"
}

func main() {
	animales := []interface{}{
		123, Canguro{}, "un gato", Perro{}, Rana{},
	}

	for _, a := range animales {
		fmt.Printf("%#v", a)
		if p, ok := a.(Perro); ok {
			fmt.Print(": ", p.Ladra())
		}
		if r, ok := a.(Rana); ok {
			fmt.Print(": ", r.Croa())
		}
		if c, ok := a.(Canguro); ok {
			fmt.Print(": ", c.Salta())
		}
		fmt.Println()
	}
}
