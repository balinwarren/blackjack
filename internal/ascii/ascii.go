package ascii

import (
	"fmt"
)

type Card struct {
	Art   []string
	Value []int
}

var title = `                                    		 .------.
    ____  __           __     _            __    |A.--. |
   / __ )/ /___ ______/ /__  (_)___ ______/ /__  | (\.------.
  / __  / / __ ` + "`" + `/ ___/ //_/ / / __ ` + "`" + `/ ___/ //_/  | :\|J.--. |
 / /_/ / / /_/ / /__/ ,<   / / /_/ / /__/ ,<     | '-| :(): |
/_____/_/\__,_/\___/_/|_|_/ /\__,_/\___/_/|_|    ` + "`" + `---| ()() |
                       /___/                         | '--'J|
                                 ` + "		     `" + `------'
    `

var ace = Card{
	Art: []string{
		".------.",
		"|A.--. |",
		`| (\/) |`,
		`| :\/: |`,
		`| '--'A|`,
		"`------'",
	},
	Value: []int{1, 11},
}

var two = Card{
	Art: []string{
		".------.",
		"|2.--. |",
		`| (\/) |`,
		`| :\/: |`,
		`| '--'2|`,
		"`------'",
	},
	Value: []int{2},
}

var three = Card{
	Art: []string{
		".------.",
		"|3.--. |",
		`| :(): |`,
		`| ()() |`,
		`| '--'3|`,
		"`------'",
	},
	Value: []int{3},
}

var four = Card{
	Art: []string{
		".------.",
		"|4.--. |",
		`| :/\: |`,
		`| :\/: |`,
		`| '--'4|`,
		"`------'",
	},
	Value: []int{4},
}

var five = Card{
	Art: []string{
		".------.",
		"|5.--. |",
		`| :/\: |`,
		`| (__) |`,
		`| '--'5|`,
		"`------'",
	},
	Value: []int{5},
}

var six = Card{
	Art: []string{
		".------.",
		"|6.--. |",
		`| (\/) |`,
		`| :\/: |`,
		`| '--'6|`,
		"`------'",
	},
	Value: []int{6},
}

var seven = Card{
	Art: []string{
		".------.",
		"|7.--. |",
		`| :(): |`,
		`| ()() |`,
		`| '--'7|`,
		"`------'",
	},
	Value: []int{7},
}

var eight = Card{
	Art: []string{
		".------.",
		"|8.--. |",
		`| :/\: |`,
		`| :\/: |`,
		`| '--'8|`,
		"`------'",
	},
	Value: []int{8},
}

var nine = Card{
	Art: []string{
		".------.",
		"|9.--. |",
		`| :/\: |`,
		`| (__) |`,
		`| '--'9|`,
		"`------'",
	},
	Value: []int{9},
}

var ten = Card{
	Art: []string{
		".------.",
		"|10--. |",
		`| (\/) |`,
		`| :\/: |`,
		`| '--10|`,
		"`------'",
	},
	Value: []int{10},
}

var jack = Card{
	Art: []string{
		".------.",
		"|J.--. |",
		`| :(): |`,
		`| ()() |`,
		`| '--'J|`,
		"`------'",
	},
	Value: []int{10},
}

var queen = Card{
	Art: []string{
		".------.",
		"|Q.--. |",
		`| (\/) |`,
		`| :\/: |`,
		`| '--'Q|`,
		"`------'",
	},
	Value: []int{10},
}

var king = Card{
	Art: []string{
		".------.",
		"|K.--. |",
		`| :/\: |`,
		`| :\/: |`,
		`| '--'K|`,
		"`------'",
	},
	Value: []int{10},
}

var hidden = []string{
	".------.",
	"|******|",
	`|**/\**|`,
	`|*(__)*|`,
	`|*'--'*|`,
	"`------'",
}

var deck = []Card{ace, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king}

func GetTitle() string {
	return title
}

func GetDeck() []Card {
	return deck
}

func GetHidden() []string {
	return hidden
}

func PrintHand(hand []Card) {
	spacer := "   "

	for i := 0; i < len(hand[0].Art); i++ {
		for j := 0; j < len(hand); j++ {
			fmt.Print(hand[j].Art[i] + spacer)
		}
		fmt.Println()
	}
}
