package main

import "fmt"

// Sin argumentos, sin valor de retorno
func DecirHola() {
	fmt.Println("hola")
}

// multiples argumentos, un solo tipo de retorno
func CoordenadaCelda(row rune, col int) string {
	return fmt.Sprint(row, col)
}

// Agrupando argumentos del mismo tipo
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Múltiples valores de retorno
func MaxMin(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

// Valores de retorno con nombre
func CuentaMayusMinus(text string) (mayus int, minus int) {
	mayus = 0
	minus = 0
	for _, c := range text {
		if c >= 'A' && c <= 'Z' {
			mayus++
		} else if c >= 'a' && c <= 'z' {
			minus++
		}
	}
	return
}

// Funcion con argumentos variables
func Sumatorio(n ...int) int {
	sum := 0
	for _, s := range n {
		sum += s
	}
	return sum
}

func main() {
	DecirHola()
	fmt.Println("La celda superior izquierda es ", CoordenadaCelda('A', 1))
	fmt.Println("El máximo de 5 y 6 es", Max(5, 6))

	ma, mi := MaxMin(4, 7)
	fmt.Printf("%v es mayor que %v\n", ma, mi)

	str := "Yo me llamo Ralph"
	up, lo := CuentaMayusMinus(str)
	fmt.Printf("%q tiene %d mayúsculas y %d minúsculas\n", str, up, lo)

	fmt.Println("3+4+5+6 =", Sumatorio(3, 4, 5, 6))

	nums := []int{7, 8, 3}
	all := Sumatorio(nums...) // operador difusor (ver capítulo 6)

	fmt.Printf("Sumando todos los números de %v: %v\n", nums, all)
}
