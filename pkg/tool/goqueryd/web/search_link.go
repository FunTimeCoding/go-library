package web

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func searchLink(
	query string,
	collection string,
	metadata map[string]string,
) string {
	parts := []string{fmt.Sprintf("query=%s", url.QueryEscape(query))}

	if collection != "" {
		parts = append(
			parts,
			fmt.Sprintf("collection=%s", url.QueryEscape(collection)),
		)
	}

	keys := make([]string, 0, len(metadata))

	for k := range metadata {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		parts = append(
			parts,
			fmt.Sprintf(
				"metadata[%s]=%s",
				url.QueryEscape(k),
				url.QueryEscape(metadata[k]),
			),
		)
	}

	return fmt.Sprintf("/search?%s", strings.Join(parts, "&"))
}
