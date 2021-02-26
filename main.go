package main

import "fmt"

func main() {

	deck := generateDeck()
	var card, nextCard card
	var guess string

	fmt.Println("Welcome to a game of High or Low\nPlease wait while the dealer shuffles the cards ... \n")
	deck.shuffle()

	fmt.Println("\nRound: ", 1)
	card, deck = deck.dealCard()
	printCard(card)

	for i := 1; i < len(deck)-1; i++ {

		nextCard, deck = deck.dealCard()
		fmt.Scanf("%v", &guess)

		if guess == "higher" {
			if nextCard.value > card.value {
				fmt.Println("Well done...")
				printCard(nextCard)

			} else {
				fmt.Println("Unlucky...")
				printCard(nextCard)
				break
			}
		} else {
			if nextCard.value < card.value {
				fmt.Println("Well done...")
				printCard(nextCard)
			} else {
				fmt.Println("Unlucky...")
				printCard(nextCard)
				break
			}
		}

		nextCard = card
		if len(deck) == 2 {
			break
		}
	}
}

func printCard(c card) {
	fmt.Println("Your card is: ", c.card)
}
