package game

import (
	"math/rand/v2"

	"github.com/balinwarren/blackjack/internal/ascii"
)

func GenerateHand() [][]string {
	deck := ascii.GetDeck()
	return [][]string{deck[rand.IntN(14)], deck[rand.IntN(14)]}
}
