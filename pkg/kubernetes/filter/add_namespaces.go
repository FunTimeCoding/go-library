package filter

func (f *Filter) AddNamespaces(s ...string) *Filter {
	f.Namespaces = append(f.Namespaces, s...)

	return f
}
