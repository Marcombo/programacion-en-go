package main

import "log"

func main() {
	log.Println("Empezando el programa")
	for i := 1; i <= 3; i++ {
		log.Printf("Mostrando evento %d\n", i)
	}
	log.Println("Saliendo del programa")
}
