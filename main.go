package main

func main() {
	// Instantiate new deck of cards
	cards := newDeck()

	// Deal returns two values, so we assign the first returned value
	// to hand, and the second to remaining cards
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
