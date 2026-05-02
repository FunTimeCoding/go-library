package option

func (f *Format) Copy() *Format {
	return new(*f)
}
