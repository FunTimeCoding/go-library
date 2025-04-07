package status

import (
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag_line"
)

type Status struct {
	bubbles    []string
	lines      []string
	linesByTag []*tag_line.Line

	format *option.Format
}
