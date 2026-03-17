package anthropic

import (
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.TokenEnvironment))
}
