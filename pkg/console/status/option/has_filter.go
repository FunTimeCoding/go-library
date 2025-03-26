package option

func (f *Format) HasFilter(
	k string,
	v string,
) bool {
	for _, p := range f.Filters {
		if p.Match(k, v) {
			return true
		}
	}

	return false
}
