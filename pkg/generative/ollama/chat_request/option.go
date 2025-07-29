package chat_request

func (r *Request) Option(
	k string,
	v any,
) *Request {
	r.request.Options[k] = v

	return r
}
