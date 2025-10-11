package locator

func (l *Locator) Trail() *Locator {
	l.trail = true

	return l
}
