package locator

func (l *Locator) Path(s string) *Locator {
	l.path = s

	return l
}
