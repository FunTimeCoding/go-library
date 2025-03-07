package merge_request

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/format"
)

func (r *Request) formatAge(s *format.Settings) string {
	result := units.HumanDuration(r.Age())

	if s.UseColor && r.AgeColor != nil {
		result = r.AgeColor(result)
	}

	return result
}
