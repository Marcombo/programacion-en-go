package main

import "fmt"

func main() {
	gastos := map[string]int{
		"Hipoteca":               812,
		"Comida":                 200,
		"Internet y telefono":    100,
		"Transporte":             80,
		"Suministros de la casa": 145,
		"Salir":                  150,
		"Otros":                  350,
	}
	fmt.Println("Tus gastos mensuales")
	total := 0
	for clave, valor := range gastos {
		fmt.Printf(" - %s: %d\n", clave, valor)
		total += valor
	}

	fmt.Printf("\nTotal mensual: %d\n", total)

	fmt.Println("\nListado de conceptos:\n")
	// Iterando solo las claves
	for clave := range gastos {
		fmt.Println(clave)
	}
}
