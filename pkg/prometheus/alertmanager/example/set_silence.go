package example

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func SetSilence() {
	alertmanager.NewEnvironment().SimpleSilence(constant.NodeNotReady)
}
