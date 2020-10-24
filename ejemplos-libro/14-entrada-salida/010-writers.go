package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := os.Stdout.Write([]byte("hola!\n"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("escritos", n, "bytes")
}
