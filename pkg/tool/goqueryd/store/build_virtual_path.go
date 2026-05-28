package store

import (
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
)

func buildVirtualPath(
	collection string,
	path string,
) string {
	return join.Empty("qmd://", collection, separator.Slash, path)
}
