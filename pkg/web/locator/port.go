package locator

func (l *Locator) Port(p int) *Locator {
	l.port = p

	return l
}
