package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func FilterAlerts(v []*Alert) []*Alert {
	var result []*Alert

	for _, element := range v {
		if element.State == constant.SuppressedState {
			continue
		}

		if element.Severity == constant.InfoSeverity {
			continue
		}

		result = append(result, element)
	}

	return result
}
