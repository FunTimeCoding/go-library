package rack

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (r *Rack) status(f *option.Format) string {
	var result string

	if r.Raw.Status != nil {
		if *r.Raw.Status.Label == ActiveLabel &&
			*r.Raw.Status.Value == Active {
			result = Active
		} else if *r.Raw.Status.Label == DeprecatedLabel &&
			*r.Raw.Status.Value == Deprecated {
			result = Deprecated

			if f.UseColor {
				result = console.Yellow(result)
			}
		} else {
			result = fmt.Sprintf(
				"%s (%s=%s)",
				Unexpected,
				*r.Raw.Status.Label,
				*r.Raw.Status.Value,
			)

			if f.UseColor {
				result = console.Red(result)
			}
		}
	} else {
		result = Unknown

		if f.UseColor {
			result = console.Red(result)
		}
	}

	return result
}
