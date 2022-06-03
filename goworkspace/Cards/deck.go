package main

import "fmt"

type singleCard string
type deck []string

func getNewDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (this deck) print() {
	printHeader("Printing the Current Deck of Cards !")

	for i, card := range this {
		fmt.Println(i, card)
	}

	printFooter()
}

func getNewCard() singleCard {
	return "Five of Dimonds"
}

func (this singleCard) print() {
	printHeader("Printing the Current Card !")

	fmt.Println(this)

	printFooter()
}

// func printHeader(headerTitle string) {
// 	fmt.Println("=======================================================")
// 	fmt.Println(headerTitle)
// 	fmt.Println("=======================================================")
// }

// func printFooter() {
// 	fmt.Println("-------------------------------------------------------")
// 	fmt.Println()
// }
