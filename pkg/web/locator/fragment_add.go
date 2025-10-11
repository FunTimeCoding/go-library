package locator

func (l *Locator) FragmentAdd(
	k string,
	v string,
) *Locator {
	l.fragmentValue.Add(k, v)

	return l
}
