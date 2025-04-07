package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (a *Alert) formatInstance() string {
	i := a.Detail(constant.InstanceLabel)

	if i == "" {
		return ""
	}

	if a.instance != nil {
		return a.instance(i)
	}

	return i
}
