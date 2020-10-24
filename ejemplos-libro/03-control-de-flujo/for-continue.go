package main

import "fmt"

func main() {
	for {
		fmt.Print("Salir? (s/n): ")
		var c rune
		fmt.Scanf("%c\n", &c)
		if c == 'N' || c == 'n' {
			continue
		}
		if c == 'S' || c == 's' {
			break
		}
		fmt.Println("carácter no reconocido")
	}
	fmt.Println("adiós!")
}
