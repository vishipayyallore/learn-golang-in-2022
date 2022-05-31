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

func getNewCard() string {
	return "Five of Dimonds"
}

func (this deck) print() {

	for i, card := range this {
		fmt.Println(i, card)
	}

}

func (this singleCard) print() {
	fmt.Println(this)
}
