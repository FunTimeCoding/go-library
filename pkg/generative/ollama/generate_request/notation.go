package generate_request

import (
	"github.com/funtimecoding/go-library/pkg/generative/ollama/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (r *Request) Notation() *Request {
	r.request.Format = notation.Marshal(constant.NotationFormat)

	return r
}
