package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	libraryConstant "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
)

func TestAlert(t *testing.T) {
	actual := New(
		&models.GettableAlert{
			Fingerprint: ptr.To(strings.Alfa),
			Status:      &models.AlertStatus{State: ptr.To(constant.ActiveState)},
			StartsAt:    ptr.To(strfmt.NewDateTime()),
		},
		strings.Bravo,
	)
	actual.Raw = nil
	assert.Any(
		t,
		&Alert{
			MonitorIdentifier: "alert-Alfa",
			Fingerprint:       "Alfa",
			Name:              "none",
			State:             "active",
			Severity:          "none",
			Summary:           "none",
			Message:           "none",
			Prometheus:        "none",
			Start:             ptr.To(libraryConstant.StartOfTime),

			Link: "https://Bravo/#/alerts?filter=%7Balertname%3D%22none%22%7D",
		},
		actual,
	)
}
