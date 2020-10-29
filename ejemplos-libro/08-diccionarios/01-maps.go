package main

import "fmt"

func main() {
	jar := make(map[string]string)
	jar["jur"] = "jor"

	// definición literal
	capitales := map[string]string{
		"Reino Unido": "Londres",
		"España":      "Madrid",
		"Japón":       "Tokyo",
		"Francia":     "París",
	}

	fmt.Println("La capital de Francia es", capitales["Francia"])
	capitales["Marruecos"] = "Rabat"

	// iteración
	for pais, capital := range capitales {
		fmt.Printf("La capital de %s es %s\n", pais, capital)
	}

	// comprobar si un valor está en el mapa
	pais := "Narnia"
	if capital, ok := capitales[pais]; ok {
		fmt.Println("La capital de", pais, "es", capital)
	} else {
		fmt.Println("No he encontrado ninguna capital para", pais)
	}

	delete(capitales, "Reino Unido")
}
