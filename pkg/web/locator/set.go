package locator

func (l *Locator) Set(
	k string,
	v string,
) *Locator {
	l.values.Set(k, v)

	return l
}
