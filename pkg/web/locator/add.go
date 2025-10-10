package locator

func (l *Locator) Add(
	k string,
	v string,
) *Locator {
	l.values.Add(k, v)

	return l
}
