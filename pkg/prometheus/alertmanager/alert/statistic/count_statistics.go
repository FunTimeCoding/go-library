package statistic

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"strings"
)

func Count(value []*alert.Alert) Statistic {
	var result Statistic

	for _, element := range value {
		result.Group.All++

		switch element.Severity {
		case constant.CriticalSeverity:
			result.Severity.Critical++
		case constant.WarningSeverity:
			result.Severity.Warning++
		case constant.InfoSeverity:
			result.Severity.Info++
		case constant.NoneSeverity:
			result.Severity.None++
		default:
			result.Severity.Unknown++
		}

		switch element.State {
		case constant.ActiveState:
			result.State.Active++
		case constant.SuppressedState:
			result.State.Suppressed++
		}

		if strings.HasPrefix(element.Name, "Kube") {
			result.Group.Kubernetes++
		} else {
			result.Group.Other++
		}
	}

	return result
}
