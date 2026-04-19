package tag

func NewSlice(v []response) []*Tag {
	var result []*Tag

	for i := range v {
		result = append(result, New(&v[i]))
	}

	return result
}
