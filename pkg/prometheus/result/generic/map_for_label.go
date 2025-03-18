package generic

import "github.com/funtimecoding/go-library/pkg/prometheus/result/metric"

func MapForLabel(
	r []*Result,
	label string,
) map[string][]*Result {
	result := make(map[string][]*Result)

	for _, element := range r {
		value := metric.New(element.Metric).Label(label)
		result[value] = append(result[value], element)
	}

	return result
}
