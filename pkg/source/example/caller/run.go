package caller

import "github.com/funtimecoding/go-library/pkg/source/example"

func Run(name string) string {
	if example.IsValid(name) {
		return example.Check(name)
	}

	return ""
}
