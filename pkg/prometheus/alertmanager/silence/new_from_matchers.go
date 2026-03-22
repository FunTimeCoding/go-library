package silence

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func NewFromMatchers(
	start *time.Time,
	end *time.Time,
	m []*models.Matcher,
) *Silence {
	return &Silence{
		Start: start,
		End:   end,
		Raw:   &models.GettableSilence{Silence: models.Silence{Matchers: m}},
	}
}
