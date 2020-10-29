package main

import (
	"fmt"
	"reflect"
)

func main() {
	original := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(original))

	n := copy(copia, original)
	fmt.Println(n, "numeros copiados:", copia)

	str := "😅ola"
	a := str[:]
	fmt.Println(a, reflect.TypeOf(a))
}
