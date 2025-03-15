package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	prometheusConstant "github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.HostEnvironment, 1),
		environment.Get(prometheusConstant.UserEnvironment, 1),
		environment.Get(prometheusConstant.PasswordEnvironment, 1),
		prometheus.NewEnvironment(),
	)
}
