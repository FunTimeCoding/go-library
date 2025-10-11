package locator

func (l *Locator) Set(
	k string,
	v string,
) *Locator {
	l.value.Set(k, v)

	return l
}
