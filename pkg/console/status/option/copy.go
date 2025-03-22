package option

func (f *Format) Copy() *Format {
	other := *f

	return &other
}
