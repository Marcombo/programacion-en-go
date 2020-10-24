package main

import "fmt"

func main() {

	var c rune
	for c != 'S' && c != 's' {
		fmt.Print("Salir? (s/n): ")
		fmt.Scanf("%c\n", &c)
	}
	fmt.Println("adiós!")
}
