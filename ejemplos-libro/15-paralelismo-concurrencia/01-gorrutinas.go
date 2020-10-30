package main

import "fmt"

func cincoVeces(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("(%d de 5) %s\n", i, msg)
	}
}

func main() {
	fmt.Println("Lanzando gorutina")

	go cincoVeces("Esta gorutina no siempre se completará")
	cincoVeces("Este mensaje se mostrará exactamente 5 veces")

	fmt.Println("Finalizando programa")
}
