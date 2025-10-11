package locator

func (l *Locator) User(u string) *Locator {
	l.user = u

	return l
}
