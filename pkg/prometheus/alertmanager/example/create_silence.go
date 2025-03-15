package example

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func CreateSilence() {
	// TODO: If not expired, extend silence by 10 minutes if already silenced
	alertmanager.NewEnvironment().SimpleSilence(constant.NodeNotReady)
}
