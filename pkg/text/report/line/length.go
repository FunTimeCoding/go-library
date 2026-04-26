package line

func (l *Line) Length() int {
	return len(l.Render())
}
