package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.Buffer{}
	buff.WriteString("hello\n")
	str, _ := buff.ReadString('\n') // ignorando el error, por simplicidad
	fmt.Println("leido:", str)

	fmt.Fprintf(&buff, "donuts %.02f\n", 12.0)

	var clave string
	var valor float32
	fmt.Fscanf(&buff, "%s %f", &clave, &valor)
	fmt.Printf("clave: %v. valor: %v\n", clave, valor)

}
