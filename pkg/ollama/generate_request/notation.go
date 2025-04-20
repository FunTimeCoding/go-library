package generate_request

import "github.com/funtimecoding/go-library/pkg/notation"

func (r *Request) Notation() *Request {
	r.request.Format = notation.Marshall("json")

	return r
}
