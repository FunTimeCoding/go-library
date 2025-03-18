package alert

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"slices"
)

func FilterSevere(v []*Alert) []*Alert {
	var result []*Alert

	for _, element := range v {
		if element.State != constant.ActiveState {
			continue
		}

		if !slices.Contains(constant.SevereSeverities, element.Severity) {
			continue
		}

		result = append(result, element)
	}

	return result
}
