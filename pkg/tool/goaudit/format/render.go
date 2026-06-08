package format

import "strings"

func (t *table) render() string {
	var b strings.Builder

	for i, c := range t.columns {
		if i > 0 {
			b.WriteString("  ")
		}

		b.WriteString(pad(c.header, c.width))
	}

	b.WriteByte('\n')

	for _, row := range t.rows {
		for i, v := range row {
			if i > 0 {
				b.WriteString("  ")
			}

			if i < len(t.columns) {
				b.WriteString(pad(v, t.columns[i].width))
			}
		}

		b.WriteByte('\n')
	}

	return b.String()
}
