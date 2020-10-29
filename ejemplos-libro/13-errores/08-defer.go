package main

import "fmt"

func main() {
	fmt.Println("hola")
	defer fmt.Println("ejecución aplazada")
	fmt.Println("adiós!")
}
