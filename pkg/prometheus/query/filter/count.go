package filter

func (f *Filter) Count() int {
	return len(f.equal) + len(f.like)
}
