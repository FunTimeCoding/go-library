package generate_request

import "github.com/ollama/ollama/api"

func New() *Request {
	return &Request{request: &api.GenerateRequest{}}
}
