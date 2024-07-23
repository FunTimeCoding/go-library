package format

import "github.com/funtimecoding/go-library/pkg/system/environment"

func (s *Settings) Compact() *Settings {
	s.UseCompact = true

	environment.Get()

	return s
}
