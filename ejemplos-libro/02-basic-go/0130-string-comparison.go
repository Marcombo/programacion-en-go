package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const correctResponse = "go build"

func main() {

	fmt.Println("Golang trivia")
	fmt.Println("=============")

	fmt.Print("Question 1: which command do you use to compile a go program? ")

	keyboard := bufio.NewReader(os.Stdin)
	response, _ := keyboard.ReadString('\n') // reading line from keyboard
	response = strings.Trim(response, " \n") // trimming spaces and new line

	// You can compare strings with ==
	if response == correctResponse {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost :(")
	}

}
