package statistic

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"strings"
)

func (s *Statistic) Count(v []*alert.Alert) *Statistic {
	result := New()

	for _, a := range v {
		result.Total++
		result.Group.All++

		switch a.Severity {
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

		switch a.State {
		case constant.ActiveState:
			result.State.Active++
		case constant.SuppressedState:
			result.State.Suppressed++
		}

		if strings.HasPrefix(a.Name, constant.KubernetesPrefix) {
			result.Group.Kubernetes++
		} else {
			result.Group.Other++
		}
	}

	return result
}
