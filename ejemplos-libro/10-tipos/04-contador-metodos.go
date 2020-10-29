package main

import "fmt"

type Contador int

func (c *Contador) Incrementa() {
	*c++
}

func (c *Contador) Reinicia(nuevoValor int) {
	*c = Contador(nuevoValor)
}

func main() {
	var c Contador
	c.Incrementa()
	c.Incrementa()
	c.Incrementa()
	fmt.Println("valor:", c)
	c.Reinicia(77)
	fmt.Println("tras reinicio:", c)
}
