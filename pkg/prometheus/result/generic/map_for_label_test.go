package generic

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"testing"
)

func TestMapForLabel(t *testing.T) {
	assert.Any(
		t,
		map[string][]*Result{
			"differing-container1": {
				{
					Type:   constant.Scalar,
					Metric: `node_load1{container="differing-container1"}`,
					Value:  "1",
				},
			},
		},
		MapForLabel(
			[]*Result{
				{
					Type:   constant.Scalar,
					Metric: `node_load1{container="differing-container1"}`,
					Value:  "1",
				},
			},
			"container",
		),
	)
}
