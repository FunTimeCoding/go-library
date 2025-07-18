package alert

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (a *Alert) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Fingerprint) {
		s.String(a.Fingerprint)
	}

	if v := a.formatEntity(f); v != "" {
		s.String(v)
	}

	if f.HasTag(tag.Category) {
		if v := a.formatCategory(f); v != "" {
			s.String(v)
		}
	}

	if a.Entity == "" && a.Category == "" {
		s.String(a.formatName(f))
	}

	s.String(a.formatSeverity(f))

	if f.HasTag(tag.Instance) {
		if v := a.formatInstance(); v != "" {
			s.String(v)
		}
	}

	if a.Start != nil {
		if false {
			s.String(a.Start.Format(time.DateMinute))
		}

		s.String(fmt.Sprintf("%s ago", units.HumanDuration(a.Age())))
	}

	s.TagLine(tag.Link, "  %s", a.Link)

	if a.Runbook != constant.None {
		s.TagLine(tag.Runbook, "  Runbook: %s", a.Runbook)
	}

	if f.ShowExtended {
		if a.Summary != constant.None {
			s.Line("  Summary: %s", a.Summary)
		}

		if a.Message != constant.None {
			s.Line("  Message: %s", a.Message)
		}

		if v := a.formatRemainingLabels(f); v != "" {
			s.Line("  Labels: %s", v)
		}

		if len(a.Receivers) > 0 {
			s.Line("  Receivers: %s", join.Comma(a.Receivers))
		}
	}

	s.RawList(a.Raw)

	return s.Format()
}
