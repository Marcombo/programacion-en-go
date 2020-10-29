// Package hola implementa maneras de saludar
package hola

import "fmt"

// ConNombre retorna un efusivo saludo, dado un nombre
func ConNombre(nombre string) string {
	return fmt.Sprintf("¡Hola, %s!", nombre)
}
