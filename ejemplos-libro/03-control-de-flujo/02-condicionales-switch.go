package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Print("La arquitectura de su procesador es ")

	arch := runtime.GOARCH
	switch arch {
	case "386":
		fmt.Println("32-bit x86")
	case "amd64":
		fmt.Println("64-bit x86")
	default:
		fmt.Println(arch)
	}

	fmt.Print("Su sistema operativo es ")

	// declaración de "os" en el mismo switch
	// limita el ámbito de la variable
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("GNU/Linux")
	case "hurd":
		fmt.Println("GNU/Hurd")
	default:
		fmt.Println(os)
	}
}
