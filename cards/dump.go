package main

import "fmt"

func mainDump() {
	// Variable declaration full initialization vs infering type
	var card1 string = "Ace of Spades"
	card2 := "Five of Diamonds"
	card3 := newCard()
	fmt.Println(card1, card2, card3)

	// string slice
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")

	for i, c := range cards {
		fmt.Println(i, c)
	}

	// byte slices
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))
}

func newCard() string {
	return "Five of Diamonds"
}
