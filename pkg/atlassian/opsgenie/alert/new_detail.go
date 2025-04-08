package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func NewDetail(
	v *alert.GetAlertResult,
	p *option.Alert,
	t *team.Team,
) *Alert {
	result := &Alert{
		Identifier:      v.Id,
		SmallIdentifier: strings.ToIntegerStrict(v.TinyId),
		Status:          v.Status,
		Priority:        v.Priority,
		Seen:            v.IsSeen,
		Acknowledged:    v.Acknowledged,
		CreatedAt:       v.CreatedAt,
		Snoozed:         v.Snoozed,
		SnoozedUntil:    v.SnoozedUntil,
		UpdatedAt:       v.UpdatedAt,
		Report:          v.Report,
		Responders:      v.Responders,
		Owner:           v.Owner,
		Source:          v.Source,
		Tags:            v.Tags,
		Details:         v.Details,
		Description:     v.Description,
		RawDetail:       v,
	}
	result.enrich(p, t)

	return result
}
