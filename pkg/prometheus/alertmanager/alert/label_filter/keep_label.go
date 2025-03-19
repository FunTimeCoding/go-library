package label_filter

func (f *Filter) KeepLabel(label string) {
	f.keepLabel = append(f.keepLabel, label)
}
