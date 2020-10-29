package main

import (
	"fmt"
	"time"
)

func main() {
	d := 67 * time.Second

	fmt.Println(d, "es equivalente a", d.Milliseconds(), "milisegundos")
}
