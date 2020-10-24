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
		switch x := a.(type) {
		case Rana:
			fmt.Println("Una rana:", x.Croa())
		case Canguro:
			fmt.Println("Un canguro:", x.Salta())
		case Perro:
			fmt.Println("Un Perro:", x.Ladra())
		default:
			fmt.Println("No sé qué es exactamente esto:", x)
		}
	}
}
