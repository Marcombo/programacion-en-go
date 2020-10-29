package main

import (
	"fmt"
)

func main() {

	fmt.Print("Escriba un carácter: ")
	var c int8
	fmt.Scanf("%c", &c)

	switch {
	case c >= 'A' && c <= 'Z':
		fmt.Println("Letra mayúscula")
	case c >= 'a' && c <= 'z':
		fmt.Println("Letra minúscula")
	case c >= '0' && c <= '9':
		fmt.Println("Dígito")
	default:
		fmt.Println("Ni letra ni dígito")
	}
}
