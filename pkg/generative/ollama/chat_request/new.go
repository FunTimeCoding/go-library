package chat_request

import "github.com/ollama/ollama/api"

func New() *Request {
	return &Request{
		request: &api.ChatRequest{},
	}
}
