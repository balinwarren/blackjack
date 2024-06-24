package game

import (
	"fmt"
	"math/rand/v2"

	"github.com/balinwarren/blackjack/internal/ascii"
)

var deck []ascii.Card = ascii.GetDeck()

var dealerHand []ascii.Card
var playerHand []ascii.Card

func StartGame() {
	fmt.Println(ascii.GetTitle())

	dealerHand = generateHand()
	playerHand = generateHand()

	ascii.PrintHand([]ascii.Card{dealerHand[0], ascii.GetHidden()})
	fmt.Print("\n\n\n\n\n")
	ascii.PrintHand(playerHand)

}

func generateHand() []ascii.Card {
	return []ascii.Card{deck[rand.IntN(13)], deck[rand.IntN(13)]}
}

func DrawCard() ascii.Card {
	return deck[rand.IntN(14)]
}
