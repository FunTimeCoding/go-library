package query

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/query/filter"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestFilter(t *testing.T) {
	assert.String(
		t,
		`up{container="Charlie",namespace="Bravo",node="Alfa"}`,
		Filter(
			constant.Up,
			filter.New().Node(
				strings.Alfa,
			).Namespace(strings.Bravo).Container(strings.Charlie),
		),
	)
}

func TestNullFilter(t *testing.T) {
	assert.String(t, "up", Filter(constant.Up, nil))
}

func TestSumByRate(t *testing.T) {
	assert.String(
		t,
		`sum(rate(kube_pod_container_status_restarts_total{namespace="Alfa"}[5m])) by (namespace, pod)`,
		SumByRate(
			constant.Restart,
			filter.New().Namespace(strings.Alfa),
			5,
			constant.Namespace,
			constant.Pod,
		),
	)
}

func TestFilterLike(t *testing.T) {
	assert.String(
		t,
		`up{instance=~"10.*"}`,
		Filter(constant.Up, filter.New().InstanceLike("10.*")),
	)
}

func TestChanges(t *testing.T) {
	assert.String(
		t,
		`changes(kube_pod_container_status_restarts_total{pod="Alfa"}[5m]) > 0`,
		Changes(constant.Restart, filter.New().Pod(strings.Alfa), 5),
	)
}

func TestSumByLabelReplace(t *testing.T) {
	assert.String(
		t,
		`sum by (owner) (label_replace(up{instance="Alfa"}, "owner", "$1", "namespace", "([^-]*)-.*"))`,
		SumByLabelReplace(
			Filter(constant.Up, filter.New().Instance(strings.Alfa)),
			"namespace",
			"owner",
			"([^-]*)-.*",
		),
	)
}
