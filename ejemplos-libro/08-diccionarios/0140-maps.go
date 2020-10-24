package main

import "fmt"

func main() {
	jar := make(map[string]string)

	jar["jur"] = "jor"

	// Map literals
	capitals := map[string]string{
		"United Kingdom": "London",
		"Spain":          "Madrid",
		"Japan":          "Tokyo",
		"France":         "Paris",
	}
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}

	fmt.Println("The capital of France is", capitals["France"])

	capitals["Morocco"] = "Rabat"

	country := "Narnia"
	if capital, ok := capitals[country]; ok {
		fmt.Println("The capital of", country, "is", capital)
	} else {
		fmt.Println("I don't find any registered capital for", country)
	}

	delete(capitals, "United Kingdom")
}
