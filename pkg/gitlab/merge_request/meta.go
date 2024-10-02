package merge_request

import "github.com/funtimecoding/go-library/pkg/console/description"

func (r *Request) Meta() *description.Description {
	return description.New("Merge request", "Merge request")
}
