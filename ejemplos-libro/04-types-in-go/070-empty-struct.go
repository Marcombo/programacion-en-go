package main

import "fmt"

func main() {
	duplicados := []string{
		"Juan", "María", "Benito", "Juan", "Carlos", "Mario",
		"Benito", "Carlos", "María", "Juan", "Cristina",
		"Isabel", "Carlos", "Juan", "Belinda",
	}

	unicos := map[string]struct{}{}

	for _, nombre := range duplicados {
		unicos[nombre] = struct{}{}
	}

	fmt.Println("La lista de nombres únicos es: ")
	for nombre := range unicos { // remember we can interate only the keys
		fmt.Println(" -", nombre)
	}
}
