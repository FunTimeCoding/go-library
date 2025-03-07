package merge_request

import "github.com/funtimecoding/go-library/pkg/face"

func (r *Request) SetAgeColorFunction(f face.SprintFunction) {
	r.AgeColor = f
}
