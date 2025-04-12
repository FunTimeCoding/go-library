package filter

import "slices"

func (f *Filter) ContainsCluster(name string) bool {
	return len(f.Clusters) == 0 || slices.Contains(f.Clusters, name)
}
