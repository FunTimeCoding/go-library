package environment

import "github.com/funtimecoding/go-library/pkg/strings/split"

func Slice(name string) []string {
	var result []string

	if s := Fallback(name, ""); s != "" {
		result = split.Comma(s)
	}

	return result
}
