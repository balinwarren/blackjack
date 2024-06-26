package cmd

import (
	"github.com/balinwarren/blackjack/internal/game"
	multiinput "github.com/balinwarren/blackjack/internal/ui/multiInput"
	tea "github.com/charmbracelet/bubbletea"
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
		game.StartGame()
		tprogram := tea.NewProgram(multiinput.InitialModelMulti([]string{"Hit", "Stand"}))
		if _, err := tprogram.Run(); err != nil {
			panic(err)
		}
	},
}
