package option

func (f *Format) Raw() *Format {
	f.ShowRaw = true

	return f
}
