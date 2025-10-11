package locator

func (l *Locator) UserPassword(
	u string,
	p string,
) *Locator {
	l.user = u
	l.password = p

	return l
}
