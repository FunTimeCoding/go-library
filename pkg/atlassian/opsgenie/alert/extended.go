package alert

import (
	"github.com/docker/go-units"
	opsgenieConstant "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"time"
)

func (a *Alert) extended(
	s *status.Status,
	f *option.Format,
) {
	s.Line("  %s", a.Link)

	if f.HasTag(tag.Name) {
		s.Line("  Name: %s", a.Name)
	}

	investigate := f.HasTag(tag.Investigate)

	if investigate && a.UpdatedAt.Sub(a.CreatedAt) > time.Minute {
		s.Line("  Update: %s", condenseTime(a.UpdatedAt))
	}

	if a.Snoozed {
		s.Line(
			"  Snoozed: %s",
			a.SnoozedUntil.Format(timeLibrary.DateMinute),
		)
	}

	if investigate && a.Report.AckTime > 0 {
		s.Line(
			"  AckTime: %s",
			units.HumanDuration(
				time.Duration(a.Report.AckTime/1000)*time.Second,
			),
		)
	}

	if a.Report.AcknowledgedBy != a.Owner {
		by := a.shortenUser(a.Report.AcknowledgedBy)

		if f.UseColor {
			by = console.Yellow("%s", by)
		}

		s.Line("  Acknowledged: %s", by)
	}

	if a.Report.CloseTime > 0 {
		s.Line("  Closed: %s", a.Report.ClosedBy)
	}

	dense := f.HasTag(tag.Dense)

	if !dense {
		var responders []string

		for _, r := range a.Responders {
			if r.Type == alert.TeamResponder {
				var name string

				if t := a.TeamMap.ByIdentifier(r.Id); t != nil {
					name = a.TeamMap.KeyByName(t.Name)

					if name == opsgenieConstant.NoKey {
						name = UnknownTeam
					}
				} else {
					name = UnknownTeam
				}

				if name == UnknownTeam {
					s.Line("  Unknown responder team: %+v", r)
				} else {
					responders = append(responders, name)
				}
			} else {
				var name string

				if u := a.UserMap.ByIdentifier(r.Id); u != nil {
					name = a.shortenUser(u.Name)

					if name == opsgenieConstant.NoKey {
						name = UnknownUser
					}
				} else {
					name = UnknownUser
				}

				if name == UnknownUser {
					s.Line("  Unknown responder user: %+v", r)
				} else {
					responders = append(responders, name)
				}
			}
		}

		if len(responders) > 0 {
			s.Line("  Responders: %s", join.Comma(responders))
		}

		if len(a.Tags) > 0 {
			s.Line("  Tags: %s", join.Comma(a.Tags))
		}

		if a.Source != "" {
			s.Line("  Source: %s", a.Source)
		}
	}
}
