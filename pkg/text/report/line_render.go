package report

func (l *line) Render() string {
	return spaces(l.indent) + l.value
}
