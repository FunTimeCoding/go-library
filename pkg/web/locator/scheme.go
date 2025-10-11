package locator

func (l *Locator) Scheme(s string) *Locator {
	l.scheme = s

	return l
}
