package example

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func SetSilence() {
	common.Alertmanager().SimpleSilence(constant.NodeNotReady)
}
