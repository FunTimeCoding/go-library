package statistic

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"testing"
)

func TestCountStatistics(t *testing.T) {
	assert.Any(
		t,
		&Statistic{
			Total:    1,
			Severity: SeverityCount{Critical: 1},
			State:    StateCount{Active: 1},
			Group:    GroupCount{All: 1, Other: 1},
		},
		New().Count(
			[]*alert.Alert{
				{
					State:    constant.ActiveState,
					Severity: constant.CriticalSeverity,
				},
			},
		),
	)
}
