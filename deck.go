package main

import (
	"fmt"
	"math/rand"
	"time"
)

type card struct {
	card  string
	value int
}
type myDeck []card

func generateDeck() myDeck {
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	deck := myDeck{}

	for _, suit := range cardSuits {

		counter := 1
		for _, value := range cardValues {

			card := card{value + " of " + suit, counter}
			deck = append(deck, card)

			counter++
		}
	}

	return deck
}

// function to print all cards in the deck
func (d myDeck) printDeck() {
	for _, card := range d {
		fmt.Println(card.card, card.value)
	}
}

//shuffling the deck function
func (d myDeck) shuffle() {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d myDeck) dealCard() (card, myDeck) {

	card := d[len(d)-1]
	d = d[:len(d)-1]

	return card, d
}
