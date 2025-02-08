package monitor

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea/constant"
	"github.com/funtimecoding/go-library/pkg/separator"
)

func (m *Model) View() string {
	return constant.Style.Render(m.table.View()) + separator.Unix
}
