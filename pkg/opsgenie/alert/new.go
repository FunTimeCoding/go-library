package alert

import (
	"github.com/funtimecoding/go-library/pkg/opsgenie/alert/option"
	"github.com/funtimecoding/go-library/pkg/opsgenie/team"
	stringLibrary "github.com/funtimecoding/go-library/pkg/strings"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func New(
	v *alert.Alert,
	p *option.Alert,
	t *team.Team,
) *Alert {
	result := &Alert{
		Identifier:      v.Id,
		SmallIdentifier: stringLibrary.ToIntegerStrict(v.TinyID),
		Status:          v.Status,
		Priority:        v.Priority,
		Seen:            v.Seen,
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
		Raw:             v,
	}
	result.enrich(p, t)

	return result
}
