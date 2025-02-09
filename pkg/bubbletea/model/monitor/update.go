package monitor

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/key"
	"github.com/funtimecoding/go-library/pkg/time"
	"log"
)

func (m *Model) Update(s tea.Msg) (tea.Model, tea.Cmd) {
	switch g := s.(type) {
	case tea.WindowSizeMsg:
		m.width = g.Width
		m.height = g.Height

		m.table.Columns()[0].Width = g.Width - 2

		m.table.SetHeight(g.Height - 4)
		m.table.SetWidth(g.Width - 2)
	case tea.KeyMsg:
		switch g.String() {
		case key.Escape:
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case key.Q, key.CtrlC:
			return m, tea.Quit
		case key.Enter:
			selected := m.table.SelectedRow()
			log.Printf("Selected log: %s\n", selected[0])

			return m, tea.Printf(
				"Selected tea: %s", selected[0],
			)
		}
	case TickMessage:
		if false {
			log.Printf("Tick: %s\n", g.Time.Format(time.DateMinute))
		}

		var commands tea.BatchMsg

		if m.second%60 == 0 {
			commands = append(commands, fetch())
		}

		m.bottomBar = fmt.Sprintf(
			"%s %d",
			g.Time.Format(time.DateSecond),
			m.second,
		)
		m.second++
		commands = append(commands, tick())

		return m, tea.Batch(commands...)
	case FetchMessage:
		var rows []table.Row

		for _, element := range g.Items {
			rows = append(rows, table.Row{element.Name})
		}

		m.table.SetRows(rows)
	}

	t, result := m.table.Update(s)
	m.table = &t

	return m, result
}
