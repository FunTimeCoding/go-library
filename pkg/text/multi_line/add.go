package multi_line

import "fmt"

func (m *MultiLine) Add(
	line string,
	arguments ...any,
) {
	if line == "" {
		return
	}

	if len(arguments) > 0 {
		m.lines = append(m.lines, fmt.Sprintf(line, arguments...))
	} else {
		m.lines = append(m.lines, line)
	}
}
