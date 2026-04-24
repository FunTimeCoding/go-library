package report

import "fmt"

func (l *line) Render() string {
	return fmt.Sprintf("%s%s", spaces(l.indent), l.value)
}
