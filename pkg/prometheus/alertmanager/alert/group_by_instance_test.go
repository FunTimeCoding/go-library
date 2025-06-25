package alert

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/prometheus/alertmanager/api/v2/models"
	"github.com/prometheus/common/model"
	"testing"
)

func TestGroupByInstance(t *testing.T) {
	assert.Any(
		t,
		map[string][]*Alert{
			"instance1": {
				{
					Name:        "Alfa",
					Fingerprint: "fingerprint1",
					Labels:      models.LabelSet{"instance": "instance1"},
				},
				{
					Name:        "Bravo",
					Fingerprint: "fingerprint2",
					Labels:      models.LabelSet{"instance": "instance1"},
				},
			},
			"instance2": {
				{
					Name:        "Charlie",
					Fingerprint: "fingerprint3",
					Labels:      models.LabelSet{"instance": "instance2"},
				},
			},
		},
		GroupByInstance(
			[]*Alert{
				{
					Name:        strings.Alfa,
					Fingerprint: "fingerprint1",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance1",
					},
				},
				{
					Name:        strings.Bravo,
					Fingerprint: "fingerprint2",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance1",
					},
				},
				{
					Name:        strings.Charlie,
					Fingerprint: "fingerprint3",
					Labels: models.LabelSet{
						model.InstanceLabel: "instance2",
					},
				},
			},
		),
	)
}
