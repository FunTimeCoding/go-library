package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestFilterAlerts(t *testing.T) {
	actual := FilterAlerts(
		[]*Alert{
			{
				Name:     upper.Alfa,
				State:    constant.ActiveState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     upper.Bravo,
				State:    constant.SuppressedState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     upper.Charlie,
				State:    constant.ActiveState,
				Severity: constant.InformationSeverity,
			},
		},
	)
	assert.Count(t, 1, actual)
	assert.String(t, "Alfa", actual[0].Name)
}
