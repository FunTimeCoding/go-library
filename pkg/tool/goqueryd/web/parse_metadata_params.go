package web

import (
	"net/url"
	"strings"
)

func parseMetadataParams(values url.Values) map[string]string {
	result := map[string]string{}

	for key, v := range values {
		if !strings.HasPrefix(key, "metadata[") {
			continue
		}

		if !strings.HasSuffix(key, "]") {
			continue
		}

		name := key[len("metadata[") : len(key)-1]

		if name != "" && len(v) > 0 && v[0] != "" {
			result[name] = v[0]
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
