package run

import "github.com/funtimecoding/go-library/pkg/notation"

func (r *Run) ParseNotation(a any) {
	notation.MustDecode(r.OutputString, &a, r.Verbose)
}
