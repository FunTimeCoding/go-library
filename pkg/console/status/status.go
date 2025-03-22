package status

import "github.com/funtimecoding/go-library/pkg/console/status/option"

type Status struct {
	bubbles    []string
	lines      []string
	linesByTag map[string][]string

	format *option.Format
}
