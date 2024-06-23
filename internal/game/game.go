package game

import (
	"math/rand/v2"

	"github.com/balinwarren/blackjack/internal/ascii"
)

var deck []ascii.Card = ascii.GetDeck()

func GenerateHand() []ascii.Card {
	return []ascii.Card{deck[rand.IntN(14)], deck[rand.IntN(14)]}
}

func DrawCard() ascii.Card {
	return deck[rand.IntN(14)]
}
