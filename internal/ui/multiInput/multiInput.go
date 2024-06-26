package multiinput

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.selected) == 1 {
				m.selected = make(map[int]struct{})
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			choice = selectedItemStyle.Render(choice)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = focusedStyle.Render("x")
		}

		option := focusedStyle.Render(choice)

		s += fmt.Sprintf("%s [%s] %s   ", cursor, checked, option)
	}

	return s
}
