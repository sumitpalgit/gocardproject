package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// Function should have a return type
func newCard() string {
	return "Five of Diamonds"
}

// Array is fixed array while slice is dynamic array which can be shrink and grow. both must have a return type.
