package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	//append doesn't modify the existing slice
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// Function should have a return type
func newCard() string {
	return "Five of Diamonds"
}
