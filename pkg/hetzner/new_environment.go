package hetzner

import (
	"github.com/funtimecoding/go-library/pkg/hetzner/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.TokenEnvironment))
}
