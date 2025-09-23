package gw2

import (
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.TokenEnvironment))
}
