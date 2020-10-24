package main

import "fmt"

type Comida string

const (
	ComidaPerro  = Comida("Comida de perro")
	ComidaHumano = Comida("Comida humana")
	ComidaGato   = Comida("Comida de gato")
	ComidaPajaro = Comida("Comida de pájaro")
)

type NoHambre struct{}

func (n NoHambre) Error() string {
	return "el perro no tiene hambre"
}

type NoApetecible struct {
	Ofrecido Comida
}

func (d NoApetecible) Error() string {
	return "a los perros no les gusta " + string(d.Ofrecido)
}

type Perro struct {
	TieneHambre bool
}

func (d *Perro) Alimenta(f Comida) error {
	if !d.TieneHambre {
		return NoHambre{}
	}
	if f != ComidaPerro && f != ComidaHumano {
		return NoApetecible{Ofrecido: f}
	}
	fmt.Println("comiendo", f, ": ñam ñam ñam!")
	d.TieneHambre = false
	return nil
}

func main() {
	dog := Perro{TieneHambre: true}
	food := []Comida{
		ComidaPajaro, ComidaGato, ComidaPerro, ComidaHumano,
	}
	for _, f := range food {
		err := dog.Alimenta(f)
		switch err.(type) {
		case NoApetecible:
			fmt.Println(err, "-> prueba otro tipo de comida")
		case NoHambre:
			fmt.Println(err, "-> espera unas horas")
		case error:
			fmt.Println(err, "(no esperaba esto!)")
		}
	}
}
