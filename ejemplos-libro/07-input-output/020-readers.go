package main

import (
	"fmt"
	"io"
	"os"
)

func printRead(r io.Reader) {
	slice := make([]byte, 10) // read reads until len(slice), NOT until cap(slice)

	num, err := r.Read(slice)
	if err != nil {
		fmt.Println("Reading error! ", err.Error())
		return
	}
	fmt.Println("I read:", string(slice[:num]))
}

func main() {
	fmt.Print("Escribe 10 carácteres: ")
	datos := make([]byte, 10) // read reads until len(slice), NOT until cap(slice)
	n, err := os.Stdin.Read(datos)
	if err != nil {
		fmt.Println("error leyendo:", err)
		return
	}
	fmt.Println("leídos", n, "bytes:", string(datos))
}
