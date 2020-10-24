package main

import (
	"bytes"
	"fmt"
)

func main() {
	buff := bytes.Buffer{}
	buff.WriteString("hello\n")
	str, _ := buff.ReadString('\n') // ignoring the error, for brevity
	fmt.Println("read from buffer:", str)

	fmt.Fprintf(&buff, "donuts %.02f\n", 12.0)

	var key string
	var val float32
	fmt.Fscanf(&buff, "%s %f", &key, &val)
	fmt.Printf("key: %v. value: %v\n", key, val)

}
