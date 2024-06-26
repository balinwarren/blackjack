package multiinput

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	choice   *Selection
}

type Selection struct {
	Choice string
}

func (s *Selection) Update(value string) {
	s.Choice = value
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModelMulti(choices []string, selection *Selection) model {
	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
		choice:   selection,
	}
}
