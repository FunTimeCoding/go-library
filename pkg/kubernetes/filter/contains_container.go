package filter

import "slices"

func (f *Filter) ContainsContainer(name string) bool {
	return len(f.Containers) == 0 || slices.Contains(f.Containers, name)
}
