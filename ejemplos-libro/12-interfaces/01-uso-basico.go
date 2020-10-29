package main

import "fmt"

type Saludador interface {
	Saluda() string
}

type Perro struct {
	Nombre string
}

func (d Perro) Saluda() string {
	return "Guau!!"
}

// Implementación implícita de la interfaz fmt.Stringer
func (d *Perro) String() string {
	return "Un perro llamado " + d.Nombre
}

func main() {
	var saludador Saludador = Perro{Nombre: "Bingo"}
	var str fmt.Stringer = &Perro{Nombre: "Rufo"}

	fmt.Println(saludador.Saluda())
	fmt.Println(str.String())
}
