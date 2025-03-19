package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (a *Alert) Suppressed() bool {
	return a.State == constant.SuppressedState
}
