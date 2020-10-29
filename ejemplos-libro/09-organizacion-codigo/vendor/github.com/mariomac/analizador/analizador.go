// Package analizador analiza palabras
package analizador

import (
	"fmt"

	"github.com/mariomac/sumadormapa/sumador"
)

const (
	mays = "mayúsculas"
	mins = "minusculas"
	cons = "consonantes"
	vocs = "vocales"
	desc = "desconocidos"
)

var vocales = map[rune]rune{
	'a': 'A', 'A': 'A', 'e': 'E', 'E': 'E',
	'i': 'I', 'I': 'I', 'o': 'O', 'O': 'O',
	'U': 'U', 'u': 'U',
	'á': 'A', 'Á': 'A', 'é': 'E', 'É': 'E',
	'í': 'I', 'Í': 'I', 'ó': 'O', 'Ó': 'O',
	'ú': 'U', 'Ú': 'U',
}

var vocMayus = map[rune]struct{}{
	'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	'Á': {}, 'É': {}, 'Í': {}, 'Ó': {}, 'Ú': {},
}

// Estadisticas de una palabra
func PrintEstadistica(palabra string) {
	sumas := calcula(palabra)
	fmt.Printf("La palabra %q contiene:\n", palabra)
	fmt.Printf("\t - %d mayúsculas\n", sumas[mays])
	fmt.Printf("\t - %d minúsculas\n", sumas[mins])
	fmt.Printf("\t - %d vocales\n", sumas[vocs])
	fmt.Printf("\t - %d consonantes\n", sumas[cons])
	if cd, ok := sumas[desc] ; ok && cd > 0 {
		fmt.Printf("\t - %d carácteres desconocidos\n", sumas[desc])
	}
	fmt.Println("Histograma de letras:")
	for c := 'A' ; c <= 'Z' ; c++ {
		if n, ok := sumas[string(c)] ; ok {
			fmt.Printf("\t%c : %d apariciones\n", c, n)
		}
	}
}

func calcula(palabra string) sumador.Claves {
	sumas := sumador.Claves{}
	for _, l := range palabra {
		if voc, ok := vocales[l]; ok {
			sumas.Incrementa(vocs)
			sumas.Incrementa(string(voc))
			if _, ok := vocMayus[l] ; ok {
				sumas.Incrementa(mays)
			} else {
				sumas.Incrementa(mins)
			}
		} else if l >= 'a' && l <= 'z' {
			sumas.Incrementa(mins)
			sumas.Incrementa(cons)
			sumas.Incrementa(string(l - 'a' + 'A'))
		} else if (l >= 'A' && l <= 'Z') || l == 'Ñ' {
			sumas.Incrementa(mays)
			sumas.Incrementa(cons)
			sumas.Incrementa(string(l))
		} else if l == 'ñ' {
			sumas.Incrementa(mins)
			sumas.Incrementa(cons)
			sumas.Incrementa("Ñ")
		} else {
			sumas.Incrementa(desc)
		}
	}
	return sumas
}
