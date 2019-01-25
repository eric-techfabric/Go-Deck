package main

func main() {
	// Instantiate new deck of cards
	cards := newDeckFromFile("my_cards")
	cards.print()
}
