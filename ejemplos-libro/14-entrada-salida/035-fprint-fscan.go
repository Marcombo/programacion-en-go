package main

import (
	"fmt"
	"os"
)

func main() {
	fichero, _ := os.Create("ejemplo.txt")

	n, _ := fmt.Fprintf(fichero, "Hola %v!\n", 1234)
	fichero.Close()
	fmt.Printf("escritos %v bytes\n", n)

	fichero, _ = os.Open("ejemplo.txt")

	var num int
	n, _ = fmt.Fscanf(fichero, "Hola %d", &num)
	fmt.Printf("leído %v argumento: %v\n", n, num)
	fichero.Close()
}
