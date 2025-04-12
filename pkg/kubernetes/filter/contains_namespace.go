package filter

import "slices"

func (f *Filter) ContainsNamespace(name string) bool {
	return len(f.Namespaces) == 0 || slices.Contains(f.Namespaces, name)
}
