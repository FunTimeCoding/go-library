package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func FilterAlerts(v []*Alert) []*Alert {
	var result []*Alert

	for _, e := range v {
		if e.State == constant.SuppressedState {
			continue
		}

		if e.Severity == constant.InfoSeverity {
			continue
		}

		result = append(result, e)
	}

	return result
}
