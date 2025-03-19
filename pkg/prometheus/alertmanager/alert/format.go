package alert

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (a *Alert) Format(s *format.Settings) string {
	t := status.New(s).String(
		a.formatName(s),
		a.formatSeverity(s),
	).Raw(a)

	if a.Start != nil {
		if false {
			t.String(a.Start.Format(time.DateMinute))
		}

		t.String(fmt.Sprintf("%s ago", units.HumanDuration(a.Age())))
	}

	t.TagLine(tag.Link, "  %s", a.Link)

	if a.Documentation != constant.None {
		t.TagLine(
			tag.Documentation,
			"  Documentation: %s",
			a.Documentation,
		)
	}

	if s.ShowExtended {
		if a.Summary != constant.None {
			t.Line("  Summary: %s", a.Summary)
		}

		if a.Message != constant.None {
			t.Line("  Message: %s", a.Message)
		}

		if v := a.formatRemainingLabels(s); v != "" {
			t.Line("  Labels: %s", v)
		}
	}

	return t.Format()
}
