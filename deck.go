package main

import (
	"fmt"
)

// Create a new type of 'deck'
// wivh is a slise og strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamond", "Spades", "Heards", "Clubs"}
	cardValues := []string{"Ace", "Two", "Tree", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	} 

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]

}
