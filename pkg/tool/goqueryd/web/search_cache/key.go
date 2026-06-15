package search_cache

import (
	"crypto/sha256"
	"fmt"
)

func Key(
	query string,
	collection string,
	mode string,
) string {
	h := sha256.Sum256(
		[]byte(fmt.Sprintf("%s\x00%s\x00%s", query, collection, mode)),
	)

	return fmt.Sprintf("%x", h[:8])
}
