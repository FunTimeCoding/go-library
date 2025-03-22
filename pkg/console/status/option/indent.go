package option

func (f *Format) Indent(i int) *Format {
	f.Indentation = i

	return f
}
