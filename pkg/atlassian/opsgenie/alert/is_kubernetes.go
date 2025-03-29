package alert

import "github.com/funtimecoding/go-library/pkg/prometheus/constant"

func (a *Alert) IsKubernetes() bool {
	return a.Detail(constant.Namespace) != ""
}
