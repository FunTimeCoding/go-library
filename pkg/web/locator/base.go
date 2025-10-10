package locator

func (l *Locator) Base(s string) *Locator {
	l.basePath = s

	return l
}
