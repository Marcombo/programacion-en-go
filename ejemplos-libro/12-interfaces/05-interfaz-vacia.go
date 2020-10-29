package main

import "fmt"

type Nombre struct {
	DePila   string
	Apellido string
}

func main() {
	var cualquierCosa interface{}

	cualquierCosa = 1
	cualquierCosa = true
	cualquierCosa = "hello"
	cualquierCosa = Nombre{DePila: "Alan", Apellido: "Turing"}

	fmt.Println(cualquierCosa)
}
