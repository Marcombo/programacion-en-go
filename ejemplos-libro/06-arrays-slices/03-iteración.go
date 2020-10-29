package main

import "fmt"

func main() {
	ciudades := []string{"Tokyo", "Lleida", "Paris", "Madrid"}

	for i := 0; i < len(ciudades); i++ {
		fmt.Printf("[%d] %s\n", i, ciudades[i])
	}

	for i, ciudad := range ciudades {
		fmt.Printf("[%d] %s\n", i, ciudad)
	}
}
