package page_file

import "github.com/funtimecoding/go-library/pkg/notation"

func (f *File) Encode() string {
	return notation.Encode(f, true)
}
