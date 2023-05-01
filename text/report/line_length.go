package report

func (l *line) Length() int {
	return len(l.Render())
}
