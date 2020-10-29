package main

import "fmt"

func funcion1() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recuperándonos del pánico:", p)
		}
	}()
	funcion2()
	fmt.Println("texto no mostrado")
}

func funcion2() {
	panic("ha sucedido algo realmente malo")
}

func main() {
	fmt.Println("invocando una función")
	funcion1()
	fmt.Println("saliendo con normalidad")
}
