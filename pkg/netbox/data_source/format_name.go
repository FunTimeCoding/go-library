package data_source

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (s *Source) formatName(f *option.Format) string {
	result := s.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
