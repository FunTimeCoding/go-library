package status

import "github.com/funtimecoding/go-library/pkg/console/format"

type Status struct {
	bubbles    []string
	lines      []string
	linesByTag map[string][]string

	settings *format.Settings
}
