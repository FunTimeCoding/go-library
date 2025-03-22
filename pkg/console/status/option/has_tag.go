package option

import "slices"

func (f *Format) HasTag(v string) bool {
	return slices.Contains(f.Tags, v)
}
