package main

import "fmt"

func main() {
	expenses := map[string]int{
		"Mortgage":          812,
		"Food":              200,
		"Internet & phones": 100,
		"Transportation":    80,
		"House Commodities": 145,
		"Going out":         150,
		"Other":             350,
	}
	fmt.Println("Your monthly expenses:")

	total := 0
	for key, value := range expenses {
		fmt.Printf(" - %s: %d\n", key, value)
		total += value
	}

	fmt.Printf("\nTotal of the month: %d\n", total)

	fmt.Println("\nlisting the concepts:\n")
	// we can also iterate only for the keys
	for key := range expenses {
		fmt.Println(key)
	}
}
