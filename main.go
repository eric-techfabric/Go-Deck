package main

func main() {
	// Instantiate new deck of cards
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
