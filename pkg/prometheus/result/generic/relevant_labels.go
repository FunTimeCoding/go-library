package generic

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/result/metric"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func RelevantLabels(
	v []*Result,
	ignore []string,
) []string {
	var m []*metric.Metric

	for _, element := range v {
		m = append(m, metric.New(element.Metric))
	}

	return sortLabels(
		strings.RemoveFromList(metric.DifferingLabels(m), ignore),
	)
}
