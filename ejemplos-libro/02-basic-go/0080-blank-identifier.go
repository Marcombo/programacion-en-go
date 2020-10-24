package main

// we can use the blank identifier to avoid compiler errors if we
// import an unused package
import _ "fmt"

func main() {
	// Blank identifier allows us to ignore some returned values
	a, _ := oneTwo()

	// Blank identifier also allows us avoid compiler errors on unused values
	_ = a
}