package multi_line

import "fmt"

func (l *MultiLine) Add(
	line string,
	arguments ...any,
) {
	if line == "" {
		return
	}

	if len(arguments) > 0 {
		l.lines = append(l.lines, fmt.Sprintf(line, arguments...))
	} else {
		l.lines = append(l.lines, line)
	}
}
