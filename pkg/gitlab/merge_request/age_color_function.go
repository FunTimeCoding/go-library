package merge_request

import "github.com/funtimecoding/go-library/pkg/face"

func (r *Request) AgeColorFunction() face.SprintFunction {
	return r.AgeColor
}
