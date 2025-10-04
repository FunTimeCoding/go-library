package openai

import (
	"github.com/funtimecoding/go-library/pkg/generative/openai/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.TokenEnvironment))
}
