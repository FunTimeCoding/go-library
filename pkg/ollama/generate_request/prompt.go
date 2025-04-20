package generate_request

func (r *Request) Prompt(s string) *Request {
	r.request.Prompt = s

	return r
}
