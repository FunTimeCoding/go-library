package file

import "github.com/funtimecoding/go-library/pkg/notation"

func Parse(s string) *Bookmark {
	result := &Bookmark{}
	notation.MustDecode(s, result, false)

	return result
}
