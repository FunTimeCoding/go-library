package option

func (f *Format) HasFilterKey(k string) bool {
	for _, p := range f.Filters {
		if p.Key == k {
			return true
		}
	}

	return false
}
