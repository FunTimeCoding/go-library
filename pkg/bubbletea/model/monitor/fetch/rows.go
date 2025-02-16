package fetch

import "github.com/charmbracelet/bubbles/table"

func (m *Message) Rows() []table.Row {
	var result []table.Row

	for _, i := range m.Items {
		detail := i.Detail

		if len(detail) > 50 {
			detail = detail[:50] + "..."
		}

		result = append(result, table.Row{i.Identifier, detail})
	}

	return result
}
