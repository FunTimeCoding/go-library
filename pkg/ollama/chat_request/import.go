package chat_request

func (r *Request) Import(other *Request) *Request {
	return r.Message(other.request.Messages...)
}
