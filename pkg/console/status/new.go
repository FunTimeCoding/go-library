package status

import "github.com/funtimecoding/go-library/pkg/console/status/option"

func New(f *option.Format) *Status {
	return &Status{format: f}
}
