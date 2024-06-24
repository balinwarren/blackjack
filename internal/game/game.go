package game

import (
	"math/rand/v2"

	"github.com/balinwarren/blackjack/internal/ascii"
)

var deck []ascii.Card = ascii.GetDeck()

var DealerHand []ascii.Card
var PlayerHand []ascii.Card

func StartGame() ([]ascii.Card, []ascii.Card) {
	DealerHand = GenerateHand()
	PlayerHand = GenerateHand()

	return DealerHand, PlayerHand
}

func GenerateHand() []ascii.Card {
	return []ascii.Card{deck[rand.IntN(13)], deck[rand.IntN(13)]}
}

func DrawCard() ascii.Card {
	return deck[rand.IntN(14)]
}
