package chat_request

import "github.com/ollama/ollama/api"

func (r *Request) Message(m ...api.Message) *Request {
	r.request.Messages = append(r.request.Messages, m...)

	return r
}
