package main

import "fmt"

type Aparato struct {
	Nombre string
	Precio int
}

func (g *Aparato) AplicaDescuento(d float32) {
	g.Precio = int(float32(g.Precio) * (1 - d))
}

type Telefono struct {
	Aparato
	Pulgadas int
	Bateria  int
}

func main() {
	s := Telefono{
		Aparato: Aparato{
			Nombre: "Zoomsunk 6G",
			Precio: 800,
		},
		Pulgadas: 6,
		Bateria:  2400,
	}
	// we can invoke embedded type functions
	s.AplicaDescuento(0.15)
	fmt.Printf("telefono con descuento %+v\n", s)

	// Struct embedding is not inheritance:
	// can't do that
	// var g Aparato = Telefono{}
	// nor
	// var g *Aparato = &Telefono{}
}
