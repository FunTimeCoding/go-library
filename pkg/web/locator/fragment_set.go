package locator

func (l *Locator) FragmentSet(
	k string,
	v string,
) *Locator {
	l.fragmentValue.Set(k, v)

	return l
}
