package main

import "fmt"

//var card string "We can initialize the variable outside the function but can't assign value"
// Value can be assigned any where in the function but not outside the function.

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" //Only first time := for variable initialization
	card = "Five of Diamonds"
	fmt.Println(card)
}
