package label

func (l *Label) Match(
	k string,
	v string,
) bool {
	if l.Key != k {
		return false
	}

	if l.Value != v {
		return false
	}

	return true
}
