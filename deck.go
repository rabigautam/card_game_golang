package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
func deal(cards deck, handSize int) (deck, deck) { //multiple return
	return cards[:handSize], cards[handSize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")

}

func (cards deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(cards.toString()), fs.ModeAppend.Perm())
}
 func newDeckFromFile(filename string) deck {
	fileByte, error := ioutil.ReadFile(filename)
	if error != nil { //error handleing
		fmt.Println("Error:", error)
		// return newDeck()
		// return
		os.Exit(1)
	}
	fileDataString := strings.SplitAfter(string(fileByte), ",")
	return	deck(fileDataString)
}
func (cards deck)  shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(source)
	
	for i := range cards{

		newPostion :=randomNumber.Intn(len(cards)-1)
		cards[i],cards[newPostion]=cards[newPostion],cards[i]    //swapping
	}
}