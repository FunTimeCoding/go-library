package filter

func (f *Filter) Like(
	k string,
	v string,
) *Filter {
	f.like[k] = v

	return f
}
