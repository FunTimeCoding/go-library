package space

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (s *Space) Format(f *option.Format) string {
	return status.New(f).String(s.Name).Format()
}
