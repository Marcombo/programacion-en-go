package main

import (
	"fmt"
	"runtime"
)

func main()  {

	fmt.Print("Your processor architecture is ")

	// notice there is no break statement
	arch := runtime.GOARCH
	switch arch {
	case "386":
		fmt.Println("32-bit x86")
	case "amd64":
		fmt.Println("64-bit x86")
	default:
		fmt.Println(arch)
	}

	fmt.Print("Your operating system is ")

	// initiaization + check in the same statement (limits var scope)
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
