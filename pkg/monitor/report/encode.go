package report

import "github.com/funtimecoding/go-library/pkg/notation"

func (r *Report) Encode() string {
	r.Count = len(r.Items)

	return notation.Encode(r, true)
}
