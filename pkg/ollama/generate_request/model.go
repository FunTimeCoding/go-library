package generate_request

func (r *Request) Model(s string) *Request {
	r.request.Model = s

	return r
}
