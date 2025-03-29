package alert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func FromResponse(
	v *alert.GetAlertResult,
	p *option.Alert,
	t *team.Team,
) *Alert {
	result := &Alert{
		Identifier:      v.Id,
		SmallIdentifier: stringLibrary.ToIntegerStrict(v.TinyId),
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
		RawResult:       v,
	}
	result.enrich(p, t)

	return result
}
