package filter

func (f *Filter) AddContainers(s ...string) *Filter {
	f.Containers = append(f.Containers, s...)

	return f
}
