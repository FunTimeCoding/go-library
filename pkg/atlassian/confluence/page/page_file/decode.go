package page_file

import "github.com/funtimecoding/go-library/pkg/notation"

func Decode(s string) *File {
	var result *File
	notation.DecodeStrict(s, &result, false)

	return result
}
