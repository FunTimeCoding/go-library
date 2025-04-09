package example

import (
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func SetSilence() {
	internal.Alertmanager().SimpleSilence(constant.NodeNotReady)
}
