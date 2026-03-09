package silence

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestSilence(t *testing.T) {
	actual := New(
		&models.GettableSilence{
			ID: new(strings.Alfa),
			Status: &models.SilenceStatus{
				State: new(constant.ActiveState),
			},
			Silence: models.Silence{
				CreatedBy: new(strings.Bravo),
				Comment:   new(strings.Charlie),
				StartsAt:  new(time.Scan(assert.NewMinute(0))),
				EndsAt:    new(time.Scan(assert.NewMinute(10))),
			},
		},
		strings.Delta,
	)
	actual.Start = nil
	actual.End = nil
	actual.Raw = nil
	assert.Any(
		t,
		&Silence{
			MonitorIdentifier: "silence-Alfa",
			Identifier:        "Alfa",
			State:             "active",
			Author:            "Bravo",
			Comment:           "Charlie",
			Rule:              "unknown rule",
			Link:              "https://Delta/#/silences/Alfa",
		},
		actual,
	)
}
