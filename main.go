package main

import "fmt"

func main() {

	//Intialise, delcare & assign variable deck to the deck returned from function
	deck := generateDeck()

	//Initiliase & declare card, nextCard of type card
	//Initialise & declare guess type string
	var card, nextCard card
	var guess string

	//Print welcome message
	fmt.Println("Welcome to a game of High or Low\nPlease wait while the dealer shuffles the cards ... \n")

	//use deck as receiver which implements shuffle function to randomly swap position of cards in deck
	deck.shuffle()

	//Show message indicating the first round
	fmt.Println("\nRound: ", 1)

	//Call the implented function "dealDeck()" that returns the first card and updated deck
	card, deck = deck.dealCard()
	printCard(card)

	//Iterate over the cards in the deck slice
	for i := 1; i < len(deck)-1; i++ {

		//Call the implented function "dealDeck()" that returns the first card and updated deck
		nextCard, deck = deck.dealCard()

		//use scan function to capture input from user and assign the pointer to the in memory address of value
		fmt.Scanf("%v", &guess)

		//If clause to check if user chose higher
		if guess == "higher" {

			//nested if clause to check if card.value is more than the presented card.value
			if nextCard.value > card.value {
				fmt.Println("Well done...")
				printCard(nextCard)
				continue
				//else to capture wrong answer from end user
			} else {
				fmt.Println("Unlucky...")
				printCard(nextCard)

				//The user loses, break loop
				break
			}
			//else clause to capture the value of lower from end user
		} else {

			//check if the next card value is lower than the target card
			if nextCard.value < card.value {
				fmt.Println("Well done...")
				printCard(nextCard)
				continue
				//Else statement to capture the wrong answer and break loop
			} else {
				fmt.Println("Unlucky...")
				printCard(nextCard)
				break
			}
		}

		//if the user is correct, assign the card to nextCard
		card = nextCard

		//Keep running loop until there is two cards left in the deck
		if len(deck) == 2 {
			break
		}
	}
}

//Print the card title
func printCard(c card) {
	fmt.Println("Your card is: ", c.card)
}
