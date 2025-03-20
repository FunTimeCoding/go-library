package name_filter

func (f *Filter) Keep(k string) *Filter {
	f.keep = append(f.keep, k)

	return f
}
