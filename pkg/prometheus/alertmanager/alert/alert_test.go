package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/linux/systemd/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
	"time"
)

func TestAlert(t *testing.T) {
	actual := New(
		&models.GettableAlert{
			Status:   &models.AlertStatus{State: ptr.To(constant.ActiveState)},
			StartsAt: ptr.To(strfmt.NewDateTime()),
		},
	)
	actual.Raw = nil
	assert.Any(
		t,
		&Alert{
			Name:       "none",
			State:      "active",
			Severity:   "none",
			Summary:    "none",
			Message:    "none",
			Prometheus: "none",
			Start: ptr.To(
				time.Date(
					1970,
					1,
					1,
					0,
					0,
					0,
					0,
					time.UTC,
				),
			),
		},
		actual,
	)
}
