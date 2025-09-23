package service

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (s *Service) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", s.Name)
	}

	return s.Name
}
