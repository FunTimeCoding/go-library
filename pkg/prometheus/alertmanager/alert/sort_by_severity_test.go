package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"testing"
	"time"
)

func TestSortBySeverity(t *testing.T) {
	assert.Count(
		t,
		1,
		SortBySeverity(
			[]*Alert{
				{Severity: constant.WarningSeverity, Start: new(time.Now())},
			},
			constant.Severities,
		),
	)
}
