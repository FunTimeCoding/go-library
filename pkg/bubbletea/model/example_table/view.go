package example_table

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/go-library/pkg/bubbletea/constant"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
)

func (m *Model) View() tea.View {
	return tea.NewView(constant.Table.Render(m.table.View()) + separator.Unix)
}
