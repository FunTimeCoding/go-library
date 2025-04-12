package filter

import "slices"

func (f *Filter) ContainsPod(name string) bool {
	return len(f.Pods) == 0 || slices.Contains(f.Pods, name)
}
