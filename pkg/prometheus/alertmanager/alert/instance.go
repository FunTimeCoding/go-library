package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (a *Alert) Instance() string {
	i := a.Detail(constant.InstanceLabel)

	if i != "" && a.instance != nil {
		return a.instance(i)
	}

	return i
}
