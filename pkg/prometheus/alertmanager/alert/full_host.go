package alert

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"strings"
)

func (a *Alert) FullHost() string {
	result := a.Detail(constant.InstanceLabel)
	result = strings.TrimPrefix(result, web.SecurePrefix)

	if strings.ContainsRune(result, ':') {
		return network.SplitHost(result)
	}

	return result
}
