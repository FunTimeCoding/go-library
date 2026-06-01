package convert

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

func Alerts(v []*alert.Alert) []*SlimAlert {
	var result []*SlimAlert

	for _, a := range v {
		result = append(result, Alert(a))
	}

	return result
}
