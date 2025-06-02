package fetch

import "github.com/charmbracelet/bubbles/table"

func (m *Message) Rows() []table.Row {
	var result []table.Row

	for _, i := range m.Items {
		result = append(
			result,
			table.Row{i.Identifier, i.Level, i.Detail, ""},
		)
	}

	return result
}
