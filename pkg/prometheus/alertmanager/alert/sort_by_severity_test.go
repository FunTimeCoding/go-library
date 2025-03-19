package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"testing"
	"time"
)

func TestSortBySeverity(t *testing.T) {
	now := time.Now()
	assert.Count(
		t,
		1,
		SortBySeverity(
			[]*Alert{{Severity: constant.WarningSeverity, Start: &now}},
			constant.Severities,
		),
	)
}
