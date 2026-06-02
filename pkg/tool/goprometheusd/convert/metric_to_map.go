package convert

import "github.com/prometheus/common/model"

func metricToMap(m model.Metric) map[string]string {
	result := make(map[string]string, len(m))

	for k, v := range m {
		result[string(k)] = string(v)
	}

	return result
}
