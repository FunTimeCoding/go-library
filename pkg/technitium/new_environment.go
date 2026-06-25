package technitium

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/technitium/constant"
)

func NewEnvironment() *Client {
	result := New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
	)

	if environment.Exists(constant.SelfSignedEnvironment) {
		result.SelfSigned()
	}

	return result
}
