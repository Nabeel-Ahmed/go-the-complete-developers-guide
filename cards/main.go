package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
