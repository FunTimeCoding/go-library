package generate_request

import "github.com/ollama/ollama/api"

func (r *Request) Get() *api.GenerateRequest {
	return r.request
}
