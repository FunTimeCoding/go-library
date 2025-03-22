package option

func (f *Format) Compact() *Format {
	f.UseCompact = true

	return f
}
