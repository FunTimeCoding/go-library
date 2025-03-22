package option

func (f *Format) Extended() *Format {
	f.ShowExtended = true

	return f
}
