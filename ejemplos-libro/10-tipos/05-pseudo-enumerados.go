package main

import (
	"fmt"
	"strings"
)

type PaloBaraja int

// iota always start as 0
const (
	Copas PaloBaraja = iota
	Oros
	Bastos
	Espadas
)

type Card struct {
	Number int
	Suit   PaloBaraja
}

func (f PaloBaraja) String() string {
	switch f {
	case Copas:
		return "Copas"
	case Oros:
		return "Oros"
	case Bastos:
		return "Bastos"
	case Espadas:
		return "Espadas"
	}
	return "desconocido"
}

type Mes int

// Iota se reinicia en cada grupo de constantes
// Si quieres que los enumerados empiecen en 1, solo añade
// 1 a iota
const (
	Enero Mes = iota + 1
	Febrero
	Marzo
	Abril
	Mayo
	Junio
	Julio
	Agosto
	Septiembre
	Octubre
	Noviembre
	Diciembre
)

type Flag int

// Puedes usar Iota en cualquier expresión
const (
	Importante Flag = 1 << iota
	Urgente
	Favorito
	TieneAdjunto
)

func (f Flag) String() string {
	sb := strings.Builder{}
	sb.WriteString("Indicadores: ")
	if f&Importante != 0 {
		sb.WriteString("importante ")
	}
	if f&Urgente != 0 {
		sb.WriteString("urgente ")
	}
	if f&Favorito != 0 {
		sb.WriteString("favorito ")
	}
	if f&TieneAdjunto != 0 {
		sb.WriteString("tieneAdjunto ")
	}
	return sb.String()
}

func main() {

	fmt.Println("Mi palo de la baraja favorito es", Oros)
	fmt.Println("Enero es el mes número", Enero)
	email := Importante | TieneAdjunto
	fmt.Println("Recibiste un email:", email)

	// sin embargo, nada te impide crear valores fuera de
	// los pseudo-enumerados
	v := PaloBaraja(30)
	fmt.Println("tengo una carta del palo:", v)
}
