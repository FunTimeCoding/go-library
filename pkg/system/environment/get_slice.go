package environment

import "github.com/funtimecoding/go-library/pkg/strings/split"

func GetSlice(name string) []string {
	var result []string

	if s := GetDefault(name, ""); s != "" {
		result = split.Comma(s)
	}

	return result
}
