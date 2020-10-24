package main

import (
	"fmt"
	"time"
)

func main() {
	d := 67 * time.Second

	fmt.Println(d, "is equivalent to", d.Milliseconds(), "milliseconds")
}
