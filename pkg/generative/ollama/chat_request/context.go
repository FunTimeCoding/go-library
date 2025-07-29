package chat_request

import "github.com/funtimecoding/go-library/pkg/generative/ollama/constant"

func (r *Request) Context(size int) *Request {
	return r.Option(constant.ContextSize, size)
}
