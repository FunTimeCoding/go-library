package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func New(
	v *alert.Alert,
	p *option.Alert,
	t *team.Team,
) *Alert {
	result := &Alert{
		Identifier:      v.Id,
		SmallIdentifier: strings.ToIntegerStrict(v.TinyID),
		Status:          v.Status,
		Priority:        v.Priority,
		Seen:            v.Seen,
		Acknowledged:    v.Acknowledged,
		Create:          v.CreatedAt,
		Snoozed:         v.Snoozed,
		SnoozeUntil:     v.SnoozedUntil,
		Update:          v.UpdatedAt,
		Report:          v.Report,
		Responders:      v.Responders,
		Owner:           v.Owner,
		Source:          v.Source,
		Tags:            v.Tags,
		RawList:         v,
	}
	result.enrich(p, t)

	return result
}
