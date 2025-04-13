package rule

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"testing"
)

func TestRule(t *testing.T) {
	actualAlert := NewAlert(
		&v1.AlertingRule{Name: strings.Alfa},
		&v1.RuleGroup{Name: strings.Bravo},
	)
	actualAlert.RawAlert = nil
	actualAlert.RawGroup = nil
	assert.Any(
		t,
		&Rule{
			Name:  "Alfa",
			Group: "Bravo",
		},
		actualAlert,
	)
	actualRecord := NewRecord(
		&v1.RecordingRule{Name: strings.Alfa},
		&v1.RuleGroup{Name: strings.Bravo},
	)
	actualRecord.RawRecord = nil
	actualRecord.RawGroup = nil
	assert.Any(
		t,
		&Rule{
			Name:  "Alfa",
			Group: "Bravo",
		},
		actualRecord,
	)
}
