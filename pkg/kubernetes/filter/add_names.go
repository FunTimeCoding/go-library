package filter

func (f *Filter) AddNames(s ...string) *Filter {
	f.Names = append(f.Names, s...)

	return f
}
