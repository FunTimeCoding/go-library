package filter

func (f *Filter) Equal(
	k string,
	v string,
) *Filter {
	f.equal[k] = v

	return f
}
