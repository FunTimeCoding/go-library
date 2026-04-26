package line

import "fmt"

func (l *Line) Render() string {
	return fmt.Sprintf("%s%s", spaces(l.Indent), l.Value)
}
