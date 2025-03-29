package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.UserKeyEnvironment, 1),
		environment.Get(constant.TeamKeyEnvironment, 1),
		environment.Get(constant.WebHostEnvironment, 1),
	)
}
