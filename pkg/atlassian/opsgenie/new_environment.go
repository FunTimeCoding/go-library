package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Get(constant.UserKeyEnvironment),
		environment.Get(constant.TeamKeyEnvironment),
		environment.Get(constant.WebHostEnvironment),
	)
}
