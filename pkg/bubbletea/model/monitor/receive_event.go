package monitor

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor/receive"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"log"
)

func (m *Model) receiveEvent(s receive.Message) {
	for _, line := range s.Text {
		elements := split.Comma(line)
		log.Printf("receive event: %s\n", line)
		command := elements[0]

		switch command {
		case constant.FlagAddCommand:
			handle := elements[1]
			identifier := elements[2]
			rows := m.table.Rows()
			rows[rowIndex(m.table, identifier)][2] = handle
			m.table.SetRows(rows)
			m.updateColumns()
		case constant.FlagRemoveCommand:
			handle := elements[1]
			identifier := elements[2]
			rows := m.table.Rows()
			i := rowIndex(m.table, identifier)

			if i != -1 {
				row := rows[i]
				log.Printf(
					"remove handle %s from %s\n",
					handle,
					row[0],
				)

				if row[2] == handle {
					row[2] = ""
					m.table.SetRows(rows)
					m.updateColumns()
				}
			}
		}
	}
}

func rowIndex(
	m *table.Model,
	identifier string,
) int {
	for i, r := range m.Rows() {
		if r[0] == identifier {
			return i
		}
	}

	return -1
}
