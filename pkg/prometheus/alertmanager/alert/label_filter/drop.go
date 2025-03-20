package label_filter

func (f *Filter) Drop(k string) *Filter {
	f.drop = append(f.drop, k)

	return f
}
