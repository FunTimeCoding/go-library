package filter

func (f *Filter) AddPods(s ...string) *Filter {
	f.Pods = append(f.Pods, s...)

	return f
}
