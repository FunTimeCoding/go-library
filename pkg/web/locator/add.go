package locator

func (l *Locator) Add(
	k string,
	v string,
) *Locator {
	l.value.Add(k, v)

	return l
}
