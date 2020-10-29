package main

import "fmt"

func main() {
	for {
		fmt.Print("Salir? (s/n): ")
		var c rune
		fmt.Scanf("%c\n", &c)
		if c == 'S' || c == 's' {
			break
		}
	}
	fmt.Println("adiós!")
}
