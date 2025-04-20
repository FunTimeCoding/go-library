package chat_request

import "github.com/funtimecoding/go-library/pkg/ollama/constant"

func (r *Request) User(s string) *Request {
	return r.Add(constant.UserRole, s)
}
