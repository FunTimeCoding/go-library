package store

import "github.com/funtimecoding/go-library/pkg/strings/join"

func buildVirtualPath(
	collection string,
	path string,
) string {
	return join.Empty("qmd://", collection, "/", path)
}
