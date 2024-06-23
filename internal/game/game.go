package game

import (
	"math/rand/v2"

	"github.com/balinwarren/blackjack/internal/ascii"
)

func GenerateHand() []ascii.Card {
	deck := ascii.GetDeck()
	return []ascii.Card{deck[rand.IntN(14)], deck[rand.IntN(14)]}
}
