package main

import (
	"fmt"
	"strings"
)

type CardSuit int

// iota always start as 0
const (
	Clubs CardSuit = iota
	Diamonds
	Hearts
	Spades
)

type Card struct {
	Number int
	Suit   CardSuit
}

func (f CardSuit) String() string {
	switch f {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	}
	return "unknown"
}

type Month int

// If you want your enums to start at 1, just add 1 to iota
// Iota is resetted every time the "const" word appears in the code
const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

type Flag int

// You can use iota in any expression
const (
	Important Flag = 1 << iota
	Urgent
	Favorite
	HasAttachment
)

func (f Flag) String() string {
	sb := strings.Builder{}
	sb.WriteString("Flags: ")
	if f&Important != 0 {
		sb.WriteString("important ")
	}
	if f&Urgent != 0 {
		sb.WriteString("urgent ")
	}
	if f&Favorite != 0 {
		sb.WriteString("favorite ")
	}
	if f&HasAttachment != 0 {
		sb.WriteString("hasAttachment ")
	}
	return sb.String()
}

func main() {

	fmt.Println("Mi favorite card suits is", Diamonds)
	fmt.Println("January is the mongh number", January)
	email := Important | HasAttachment
	fmt.Println("You received an email with", email)

	// however, iota does not create real enums
	v := CardSuit(30)
	fmt.Println("I have a card suited as", v)
}
