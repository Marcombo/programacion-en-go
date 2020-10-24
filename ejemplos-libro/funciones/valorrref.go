package main

import "fmt"

func Incrementa(a *int) {
	*a++
}
func main() {
	a := 3
	fmt.Println("a =", a)
	Incrementa(&a)
	fmt.Println("a =", a)
}
