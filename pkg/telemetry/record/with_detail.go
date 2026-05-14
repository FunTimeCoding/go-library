package record

func (r *Record) WithDetail(detail map[string]string) *Record {
	r.Detail = detail

	return r
}
