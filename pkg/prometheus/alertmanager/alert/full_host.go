package alert

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
)

func (a *Alert) FullHost() string {
	return network.SplitHost(a.Detail(constant.InstanceLabel))
}
