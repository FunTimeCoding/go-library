package printer

import "github.com/funtimecoding/go-library/pkg/console/format"

func New(s *format.Settings) *Printer {
	return &Printer{setting: s}
}
