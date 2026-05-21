package web

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
)

func shortenPaths(files string) string {
	var result []string

	for _, line := range split.NewLine(files) {
		result = append(result, shortenPath(line))
	}

	return join.CommaSpace(result)
}
