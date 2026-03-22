package example_table

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/bubbletea/constant"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
)

func (m *Model) View() tea.View {
	return tea.NewView(fmt.Sprintf("%s%s", constant.Table.Render(m.table.View()), separator.Unix))
}
