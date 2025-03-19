package label_filter

func (f *Filter) KeepLabel(k string) *Filter {
	f.keepLabel = append(f.keepLabel, k)

	return f
}
