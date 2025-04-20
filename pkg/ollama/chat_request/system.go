package chat_request

import "github.com/funtimecoding/go-library/pkg/ollama/constant"

func (r *Request) System(s string) *Request {
	return r.Add(constant.SystemRole, s)
}
