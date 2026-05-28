package goprocessd

import "github.com/funtimecoding/go-library/pkg/tool/goprocessd/procfile"

func filterEntries(
	entries []procfile.Entry,
	names []string,
) []procfile.Entry {
	wanted := make(map[string]bool, len(names))

	for _, name := range names {
		wanted[name] = true
	}

	var result []procfile.Entry

	for _, entry := range entries {
		if wanted[entry.Name] {
			result = append(result, entry)
		}
	}

	return result
}
