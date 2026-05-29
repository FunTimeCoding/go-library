package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"strings"
)

func sourceTypePrefixes(p string) []string {
	parts := split.Slash(strings.Trim(p, separator.Slash))
	result := make([]string, 0, len(parts))

	for i := len(parts) - 1; i >= 0; i-- {
		result = append(
			result,
			fmt.Sprintf("%s/", join.Slash(parts[:i+1])),
		)
	}

	result = append(result, "")

	return result
}
