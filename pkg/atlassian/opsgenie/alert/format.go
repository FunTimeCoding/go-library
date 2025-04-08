package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (a *Alert) Format(f *option.Format) string {
	s := status.New(f).String(a.formatName(f), a.formatPriority(f))

	if !f.HasTag(tag.Filter) && !f.HasTag(tag.Dense) {
		s.String(a.formatStatus(f))
	}

	s.String(a.formatTeamName(f), a.formatAge())

	if o := a.formatOwner(f); o != "" {
		s.String(o)
	}

	if l := a.flags(f); len(l) > 0 {
		s.String(join.Comma(l))
	}

	if f.ShowExtended {
		a.extended(s, f)
	}

	if f.ShowRaw {
		if a.RawList != nil {
			s.Raw(a.RawList)
		} else if a.RawDetail != nil {
			r := fmt.Sprintf("%+v", a.RawDetail)

			if f.UseColor {
				r = console.Magenta(r)
			}

			s.Line("  RawDetail: %s", r)
		}
	}

	return s.Format()
}
