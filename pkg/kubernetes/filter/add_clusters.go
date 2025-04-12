package filter

func (f *Filter) AddClusters(s ...string) *Filter {
	f.Clusters = append(f.Clusters, s...)

	return f
}
