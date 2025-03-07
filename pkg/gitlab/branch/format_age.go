package branch

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (b *Branch) formatAge(s *format.Settings) string {
	result := units.HumanDuration(b.Age())

	if s.UseColor && b.AgeColor != nil {
		result = b.AgeColor(result)
	}

	return result
}
