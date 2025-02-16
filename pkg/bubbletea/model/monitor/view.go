package monitor

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/funtimecoding/go-library/pkg/bubbletea/constant"
)

func (m *Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "wait for screen size"
	}

	m.topBar = fmt.Sprintf("%dx%d", m.width, m.height)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		constant.Default.Render(m.topBar),
		constant.Table.Render(m.table.View()),
		constant.Default.Render(m.bottomBar),
	)
}
