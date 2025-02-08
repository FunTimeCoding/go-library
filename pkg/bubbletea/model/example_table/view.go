package example_table

import "github.com/funtimecoding/go-library/pkg/bubbletea/constant"

func (m *Model) View() string {
	return constant.Style.Render(m.table.View()) + "\n"
}
