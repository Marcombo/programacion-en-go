package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Escribiendo archivo...")
	fichero, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if _, err = fichero.Write([]byte("hola!")); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fichero.Close() // interface Closer

	fmt.Println("Leyendo archivo... ")
	if fichero, err = os.Open("ejemplo.txt"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer fichero.Close()
	leido := make([]byte, 256)
	n, err := fichero.Read(leido)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("leídos %v bytes: %s\n", n, string(leido[:n]))

}
