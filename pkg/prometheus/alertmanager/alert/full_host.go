package alert

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"strings"
)

func (a *Alert) FullHost() string {
	result := a.Detail(constant.InstanceLabel)

	if strings.ContainsRune(result, ':') {
		return network.SplitHost(result)
	}

	return result
}
