package locator

func (l *Locator) Host(s string) *Locator {
	l.host = s

	return l
}
