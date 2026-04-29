package client

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/client/constant"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
