package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

func newModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case key.CtrlC, key.Q:
			return m, tea.Quit
		case key.Up, key.K:
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Down, key.J:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case key.Enter, key.Space:
			if _, okay := m.selected[m.cursor]; okay {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		checked := " "

		if _, okay := m.selected[i]; okay {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	bubbletea.Run(newModel())
}
