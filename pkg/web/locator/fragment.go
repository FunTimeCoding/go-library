package locator

func (l *Locator) Fragment(s string) *Locator {
	l.fragment = s

	return l
}
