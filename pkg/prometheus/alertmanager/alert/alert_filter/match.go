package alert_filter

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"time"
)

func (f *Filter) match(a *alert.Alert) bool {
	p := f.parameter

	if !p.All {
		switch a.Severity {
		case constant.NoneSeverity:
			return false
		case constant.InfoSeverity:
			return false
		}

		if !p.Old && a.Age() > 7*24*time.Hour {
			return false
		}
	}

	if p.CriticalOnly && a.Severity != constant.CriticalSeverity {
		return false
	}

	if p.WarningOnly && a.Severity != constant.WarningSeverity {
		return false
	}

	if !p.Suppressed && a.Suppressed() {
		return false
	}

	return true
}
