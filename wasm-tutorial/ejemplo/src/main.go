package main

import "syscall/js"

func main() {
	window := js.Global()
	doc := window.Get("document")
	body := doc.Get("body")
	body.Set("onclick",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			div := doc.Call("createElement", "div")
			div.Set("textContent", "¡clic!")
			body.Call("appendChild", div)
			return nil
		}))
	<-make(chan struct{})
}
