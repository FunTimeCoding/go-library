package chat_request

import "github.com/ollama/ollama/api"

func (r *Request) Get() *api.ChatRequest {
	return r.request
}
