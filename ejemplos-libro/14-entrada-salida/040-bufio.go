package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {
	buff := bytes.NewBufferString(`hola que tal
probando texto
multilínea`)

	sc := bufio.NewReader(buff)

	leido, err := sc.ReadString('\n')
	for err == nil {
		fmt.Print("leida línea:", leido)
		leido, err = sc.ReadString('\n')
	}

	if err == io.EOF {
		fmt.Println("línea final:", leido)
	} else {
		fmt.Println("error inesperado:", err.Error())
		return
	}
}
