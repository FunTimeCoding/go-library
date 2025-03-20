package statistic

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"testing"
)

func TestCountStatistics(t *testing.T) {
	s := New()
	assert.Any(
		t,
		&Statistic{
			Total: 1,
		},
		s.CountBeforeProcessing(
			[]*alert.Alert{
				{
					State:    constant.ActiveState,
					Severity: constant.CriticalSeverity,
				},
			},
		),
	)
	assert.Any(
		t,
		&Statistic{
			Total:    1,
			Relevant: 1,
			Severity: SeverityCount{Critical: 1},
			State:    StateCount{Active: 1},
			Group:    GroupCount{All: 1, Other: 1},
		},
		s.CountAfterProcessing(
			[]*alert.Alert{
				{
					State:    constant.ActiveState,
					Severity: constant.CriticalSeverity,
				},
			},
		),
	)
}
