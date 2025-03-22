package printer

import "github.com/funtimecoding/go-library/pkg/console/status/option"

func New(f *option.Format) *Printer {
	return &Printer{format: f}
}
