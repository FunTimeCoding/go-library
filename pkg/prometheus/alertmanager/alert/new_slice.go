package alert

import "github.com/prometheus/alertmanager/api/v2/models"

func NewSlice(v models.GettableAlerts) []*Alert {
	var result []*Alert

	for _, a := range v {
		result = append(result, New(a))
	}

	return result
}
