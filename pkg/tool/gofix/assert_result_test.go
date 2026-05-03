package gofix

import "testing"

func filterApplied(entries []result) []result {
	var r []result

	for _, e := range entries {
		if !e.Blocked {
			r = append(r, e)
		}
	}

	return r
}

func filterBlocked(entries []result) []result {
	var r []result

	for _, e := range entries {
		if e.Blocked {
			r = append(r, e)
		}
	}

	return r
}

func assertResult(
	t *testing.T,
	entries []result,
	path string,
	message string,
) {
	t.Helper()

	for _, e := range entries {
		if e.Path == path && e.Message == message {
			return
		}
	}

	t.Errorf(
		"expected result {path: %q, message: %q} not found",
		path,
		message,
	)
}
