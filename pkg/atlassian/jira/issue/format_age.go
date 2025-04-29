package issue

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"time"
)

func (i *Issue) FormatAge(f *option.Format) string {
	if time.Time(i.Raw.Fields.Created).IsZero() {
		return NoAge
	}

	result := units.HumanDuration(i.Age())

	if f.UseColor && i.ageColor != nil {
		result = i.ageColor(result)
	}

	return result
}
