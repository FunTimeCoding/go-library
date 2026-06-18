package caller

import "github.com/funtimecoding/go-library/pkg/source/example"

func Lookup(
	e *example.Entry,
	name string,
) bool {
	return e.FindByName(name)
}
