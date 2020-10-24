package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.NewBufferString(`hola que tal
probando escaneo
multilínea`)

	sc := bufio.NewScanner(buff)
	sc.Split(bufio.ScanLines)
	linea := 1
	for sc.Scan() {
		fmt.Printf("%d: %s\n", linea, sc.Text())
		linea++
	}
}
