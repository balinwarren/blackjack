package cmd

import (
	"fmt"

	"github.com/balinwarren/blackjack/internal/ascii"
	"github.com/balinwarren/blackjack/internal/game"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start a game of blackjack.",
	Long:  `Beat the dealer living in your terminal. Play a game of blackjack.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ascii.GetTitle())
		var hand = game.GenerateHand()
		ascii.PrintHand(hand)
	},
}
