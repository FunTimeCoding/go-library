package alert

import "github.com/prometheus/alertmanager/api/v2/models"

func NewSlice(
	v models.GettableAlerts,
	host string,
) []*Alert {
	var result []*Alert

	for _, a := range v {
		result = append(result, New(a, host))
	}

	return result
}
