package filter

import "slices"

func (f *Filter) ContainsName(name string) bool {
	return len(f.Names) == 0 || slices.Contains(f.Names, name)
}
