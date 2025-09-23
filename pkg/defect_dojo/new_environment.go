package defect_dojo

import (
	"github.com/funtimecoding/go-library/pkg/defect_dojo/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.HostEnvironment),
		environment.Required(constant.TokenEnvironment),
	)
}
