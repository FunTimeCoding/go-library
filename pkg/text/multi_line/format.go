package multi_line

import "fmt"

func (m *MultiLine) Format(
	line string,
	arguments ...any,
) {
	if len(arguments) > 0 {
		m.lines = append(m.lines, fmt.Sprintf(line, arguments...))
	} else {
		m.lines = append(m.lines, line)
	}
}
