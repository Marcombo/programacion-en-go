package main

import "fmt"

type Año int
type Nacimientos []Año

type NombreNacimientos map[string]Año

func main() {
	censo := Nacimientos{1979, 1983, 1965}
	censo = append(censo, 1990)
	censo = append(censo, 1955)

	suma := Año(0)
	for _, a := range censo {
		suma += a
	}

	media := suma / Año(len(censo))
	fmt.Println("Año de nacimiento medio:", media)

	// Same for instantiating MontlyAvgTemp
	artistas := NombreNacimientos{
		"Vincent Van Gogh": 1853,
		"Elvis Presley":    1935,
		"Salvador Dali":    1904,
	}
	artistas["Rick Astley"] = 1966
	for nombre, año := range artistas {
		fmt.Printf("%s nació en %d\n", nombre, año)
	}
}
