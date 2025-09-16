package environment

import "github.com/funtimecoding/go-library/pkg/strings/split"

func Slice(name string) []string {
	var result []string

	if s := Default(name, ""); s != "" {
		result = split.Comma(s)
	}

	return result
}
