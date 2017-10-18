package main

import "fmt"

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

	// exercise
	s := 10
	a := []int{}
	for i := 0; i <= s; i++ {
		a = append(a, i)
	}
	for i := range a {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
