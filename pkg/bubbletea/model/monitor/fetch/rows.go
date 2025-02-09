package fetch

import "github.com/charmbracelet/bubbles/table"

func (m *Message) Rows() []table.Row {
	var result []table.Row

	for _, element := range m.Items {
		result = append(result, table.Row{element.Name})
	}

	return result
}
