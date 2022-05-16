package main

// import "fmt"


func main() {

	// cards := newDeck()

	// hand, reaminingCards := deal(cards, 5)

	// hand.print()
	// reaminingCards.print()




	// cards := newDeck()
	//  // fmt.Println(cards.saveToFile())
	// cards.saveToFile("My_cards")


	// cards := newDeckFromFile("My_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
	

}

