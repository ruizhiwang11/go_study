package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // dynamic data binding, same as python
	//card = "Five of Diamonds" // change a variable
	//fmt.Println(card)
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()
	cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// hand, remaining_deck := deal(cards, 5)
	// hand.print()
	// remaining_deck.print()
	// fmt.Println(hand.toString())
	// hand.saveToFile("./123.txt")
	// f_cards := newDeckFromFile("123.txt")
	// f_cards.print()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
