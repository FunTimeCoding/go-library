package monitor

import "github.com/funtimecoding/go-library/pkg/bubbletea/constant"

func (m *Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "wait for screen size"
	}

	return constant.Style.Render(m.table.View())
}
