package option

func (f *Format) Tag(v ...string) *Format {
	f.Tags = append(f.Tags, v...)

	return f
}
