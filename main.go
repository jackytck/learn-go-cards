package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards.saveToFile("my_cards")

	fromDisk := newDeckFromFile("my_cards")
	// fromDisk.print()

	fromDisk.shuffle()
	fromDisk.print()
}
