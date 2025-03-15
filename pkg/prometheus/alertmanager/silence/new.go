package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/openapi"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func New(v *models.GettableSilence) *Silence {
	var match []string

	for _, m := range v.Matchers {
		match = append(match, fmt.Sprintf("%s=%s", *m.Name, *m.Value))
	}

	return &Silence{
		Identifier: *v.ID,
		State:      *v.Status.State,
		Match:      join.Comma(match),
		Start:      openapi.ConvertTime(v.StartsAt),
		End:        openapi.ConvertTime(v.EndsAt),
		Author:     *v.CreatedBy,
		Comment:    *v.Comment,
		Raw:        v,
	}
}
