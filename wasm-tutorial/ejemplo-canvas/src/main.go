package main

import (
	"time"

	"github.com/mariomac/gorrazo/pkg/draw"
)

func main() {
	// obtenemos el canvas con id="lienzoDibujo"
	// y lo configuramos para que ocupe toda la pantalla
	c := draw.GetCanvas("lienzoDibujo",
		draw.FullScreen(true))

	rotacion := float64(0)
	for {
		// dimensión del canvas
		w, h := c.GetWidth(), c.GetHeight()

		// creación de un gradiente de fondo
		gradient := c.CreateLinearGradient(0, 0, w, h)
		gradient.AddColorStop(1, "#ff00ff")
		gradient.AddColorStop(0.3, "#888888")
		gradient.AddColorStop(0, "#00ff88")
		c.FillStyleGradient(gradient)
		c.FillRect(0, 0, w, h)

		// Guardar el estado del canvas
		c.Save()

		// Sombra del texto
		c.ShadowBlur(15)
		c.ShadowColor("black")
		c.ShadowOffsetX(10)
		c.ShadowOffsetY(20)

		// Color y fuente del texto
		c.FillStyle("yellow")
		c.Font("40px Arial")

		// Llevar el texto hacia el centro, y rotarlo
		c.Translate(w/2, h/2)
		c.Rotate(rotacion)

		// dibujar el texto, desplazándolo 200 puntos a la izquierda
		// y 20 puntos hacia arriba
		c.FillText("¡Hola, gopher!", -200, -20)

		// restaurar el estado del canvas al que
		// era antes de llamar a c.Save()
		// eliminará sombras, rotaciones, colore....
		c.Restore()

		// actualiza rotación y espera 20 milisegundos
		rotacion += 0.02
		time.Sleep(20 * time.Millisecond)
	}
}
