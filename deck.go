package main

import "fmt"

//Create new type of deck
//slice of stringd

type deck []string

func newDeck() deck { //this function doesn't require receiver as it doesn't call method name
	cards := deck{}
	cardsSuits := []string{"Club", "Spades", "Hearts", "Diamonds"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+"of "+suit)
		}
	}
	return cards
}

func (cards deck) print() { //receiver function
	for i, card := range cards {
		fmt.Println(i, card)

	}
}
func deal(cards deck, handSize int) (deck, deck) {               //multiple return
	return cards[:handSize],cards[handSize:]
}
