package main

import (
	"fmt"
)

type Cuboide struct {
	Ancho    float64
	Alto     float64
	Profundo float64
}

func (c Cuboide) Volumen() float64 {
	return c.Ancho * c.Profundo * c.Alto
}

func (c *Cuboide) Redimensiona(an, al, pr float64) {
	c.Ancho = an
	c.Alto = al
	c.Profundo = pr
}

func (c Cuboide) String() string {
	return fmt.Sprintf("%v x %v x %v",
		c.Ancho, c.Profundo, c.Alto)
}

func main() {
	c := Cuboide{Ancho: 2, Profundo: 3, Alto: 2}
	fmt.Printf("cuboide %v. volumen %v\n", c, c.Volumen())
	c.Redimensiona(1, 2, 3)
	fmt.Printf("cuboide %v. volumen %v\n", c, c.Volumen())

	fmt.Printf("%#v\n", c)
	fmt.Println(c)
}
