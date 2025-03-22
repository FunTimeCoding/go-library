package branch

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (b *Branch) formatAge(f *option.Format) string {
	result := units.HumanDuration(b.Age())

	if f.UseColor && b.AgeColor != nil {
		result = b.AgeColor(result)
	}

	return result
}
