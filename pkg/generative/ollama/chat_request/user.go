package chat_request

import "github.com/funtimecoding/go-library/pkg/generative/ollama/constant"

func (r *Request) User(s string) *Request {
	return r.Add(constant.UserRole, s)
}
