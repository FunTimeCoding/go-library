package label_filter

func (f *Filter) DropLabel(label string) {
	f.dropLabel = append(f.dropLabel, label)
}
