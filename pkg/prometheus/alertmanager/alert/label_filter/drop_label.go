package label_filter

func (f *Filter) DropLabel(k string) *Filter {
	f.dropLabel = append(f.dropLabel, k)

	return f
}
