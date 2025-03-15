package silence

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
	"time"
)

func TestSilence(t *testing.T) {
	actual := New(
		&models.GettableSilence{
			ID: ptr.To(strings.Alfa),
			Status: &models.SilenceStatus{
				State: ptr.To(constant.ActiveState),
			},
			Silence: models.Silence{
				CreatedBy: ptr.To(strings.Bravo),
				Comment:   ptr.To(strings.Charlie),
				StartsAt: ptr.To(
					timeLibrary.Scan(
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
				),
				EndsAt: ptr.To(
					timeLibrary.Scan(
						time.Date(
							1970,
							1,
							1,
							0,
							10,
							0,
							0,
							time.UTC,
						),
					),
				),
			},
		},
	)
	actual.Start = nil
	actual.End = nil
	actual.Raw = nil
	assert.Any(
		t,
		&Silence{
			Identifier: "Alfa",
			State:      "active",
			Author:     "Bravo",
			Comment:    "Charlie",
		},
		actual,
	)
}
