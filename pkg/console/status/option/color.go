package option

func (f *Format) Color() *Format {
	f.UseColor = true

	return f
}
