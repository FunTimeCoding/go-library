package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (a *Alert) HostDetail() string {
	result := a.Detail(constant.TargetLabel)

	if result == "" {
		result = a.Detail(constant.InstanceLabel)
	}

	return result
}
