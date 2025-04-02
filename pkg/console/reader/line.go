package reader

import "github.com/funtimecoding/go-library/pkg/system"

func (r *Reader) Line() string {
	return system.ReadLine(r.reader)
}
