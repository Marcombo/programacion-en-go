// Package sumador es un ejemplo sin utilidad ninguna
package sumador

// Claves es un mapa que puede ir añadiendo números según una clave
type Claves map[string]int

// Incrementa la aparición de una clave concreta. La crea si no existe
func (c Claves) Incrementa(clave string) {
	valor, ok := c[clave]
	if ok {
		valor++
	} else {
		valor = 1
	}
	c[clave] = valor
}
