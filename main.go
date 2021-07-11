package main

// import "fmt"


func main() {
	cards := newDeck() //slice of type string
	// cards.print()
	// handDeck,remaningDeck:=deal(cards,10)
	// handDeck.print()
	// fmt.Println("above is handCard")
	// remaningDeck.print()	
	// fmt.Println(cards.toString())
	// cards.saveToFile("test go file")
	// readCards:=newDeckFromFile("test go file")
	// readCards.print()
	cards.shuffle()
	cards.print()
}

