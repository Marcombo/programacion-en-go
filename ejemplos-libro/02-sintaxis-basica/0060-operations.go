package main

import "fmt"

func main() {
	lang := 'C'

	fmt.Print("Los operadores de Go son similares")
	fmt.Printf("a los del lenguaje %c", lang)

	lang++

	fmt.Printf(", incluso a los de %c.\n", lang)

	p0 := 10.1
	v := 1.2
	t := 3.2

	p := p0 + v*t

	// Printf has a generic %v to print values
	fmt.Printf("Dado un objeto en la posición %v ", p0)
	fmt.Printf("con una velocidad de %v m/s, ", v)
	fmt.Printf("después de %v segundos ", t)
	fmt.Printf("estará situado en la posición %v\n", p)
}
