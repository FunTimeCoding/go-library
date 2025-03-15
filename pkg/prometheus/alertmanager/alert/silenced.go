package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (a *Alert) Silenced() bool {
	return a.State == constant.SuppressedState
}
