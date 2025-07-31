package fetch

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
)

func (m *Message) Rows() []table.Row {
	var result []table.Row

	for _, i := range m.Items {
		result = append(
			result,
			table.Row{
				i.Identifier,
				fmt.Sprintf("%.1f", i.Score),
				string(i.Severity),
				i.Detail,
				"",
			},
		)
	}

	return result
}
