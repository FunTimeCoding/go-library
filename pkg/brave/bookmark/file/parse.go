package file

import "github.com/funtimecoding/go-library/pkg/notation"

func Parse(s string) *Bookmark {
	result := &Bookmark{}
	notation.DecodeStrict(s, result, false)

	return result
}
