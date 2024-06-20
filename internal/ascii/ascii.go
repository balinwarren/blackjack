package ascii

var title = `                                    		 .------.
    ____  __           __     _            __    |A.--. |
   / __ )/ /___ ______/ /__  (_)___ ______/ /__  | (\.------.
  / __  / / __ ` + "`" + `/ ___/ //_/ / / __ ` + "`" + `/ ___/ //_/  | :\|J.--. |
 / /_/ / / /_/ / /__/ ,<   / / /_/ / /__/ ,<     | '-| :(): |
/_____/_/\__,_/\___/_/|_|_/ /\__,_/\___/_/|_|    ` + "`" + `---| ()() |
                       /___/                         | '--'J|
                                 ` + "		     `" + `------'
    `

var ace = `
.------.
|A.--. |
| (\/) |
| :\/: |
| '--'A|
` + "`" + `------'
`
var two = `
.------.
|2.--. |
| (\/) |
| :\/: |
| '--'2|
` + "`" + `------'
`
var three = `
.------.
|3.--. |
| :(): |
| ()() |
| '--'3|
` + "`" + `------'
`

var four = `
.------.
|4.--. |
| :/\: |
| :\/: |
| '--'4|
` + "`" + `------'
`

var five = `
.------.
|5.--. |
| :/\: |
| (__) |
| '--'5|
` + "`" + `------'
`

var six = `
.------.
|6.--. |
| (\/) |
| :\/: |
| '--'6|
` + "`" + `------'
`

var seven = `
.------.
|7.--. |
| :(): |
| ()() |
| '--'7|
` + "`" + `------'
`

var eight = `
.------.
|8.--. |
| :/\: |
| :\/: |
| '--'8|
` + "`" + `------'
`

var nine = `
.------.
|9.--. |
| :/\: |
| (__) |
| '--'9|
` + "`" + `------'
`

var ten = `
.------.
|10--. |
| (\/) |
| :\/: |
| '--10|
` + "`" + `------'
`

var jack = `
.------.
|J.--. |
| :(): |
| ()() |
| '--'J|
` + "`" + `------'
`

var queen = `
.------.
|Q.--. |
| (\/) |
| :\/: |
| '--'Q|
` + "`" + `------'
`

var king = `
.------.
|K.--. |
| :/\: |
| :\/: |
| '--'K|
` + "`" + `------'
`

var hidden = `
.------.
|******|
|**/\**|
|*(  )*|
|*'--'*|
` + "`" + `------'
`
var deck = []string{ace, two, three, four, five, six, seven, eight, nine, ten, jack, queen, king}

func GetTitle() string {
	return title
}

func GetDeck() []string {
	return deck
}

func GetHidden() string {
	return hidden
}
