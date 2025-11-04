package vault

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/vault/constant"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
