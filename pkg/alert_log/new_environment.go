package alert_log

import (
	"github.com/funtimecoding/go-library/pkg/alert_log/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
