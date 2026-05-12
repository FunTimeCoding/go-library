package testutil

import "github.com/funtimecoding/go-library/pkg/lint/output"

func countBlocked(results *output.Results) int {
	var count int

	for _, e := range results.Entries {
		if e.Blocked {
			count++
		}
	}

	return count
}
