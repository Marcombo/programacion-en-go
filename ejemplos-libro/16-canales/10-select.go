package main

import (
	"fmt"
	"time"
)

func CentralMensajeria(sms, email, carta <-chan string) {
	for {
		select {
		case num := <-sms:
			fmt.Println("recibido SMS del número", num)
		case dir := <-email:
			fmt.Println("recibido email de dirección", dir)
		case rem := <-carta:
			fmt.Println("recibida carta de remitente", rem)
		}
	}
}

func main() {
	sms := make(chan string)
	email := make(chan string)
	carta := make(chan string)

	go CentralMensajeria(sms, email, carta)

	// no se garantiza recepción ordenada entre diferentes
	// canales (sí dentro de un mismo canal)
	sms <- "777889923"
	email <- "yahoo@google.com"
	carta <- "Banco Central Hispano"
	email <- "noreply@example.com"

	time.Sleep(2 * time.Second)
}
