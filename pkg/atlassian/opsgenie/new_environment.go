package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Exit(constant.UserKeyEnvironment),
		environment.Exit(constant.TeamKeyEnvironment),
		environment.Exit(constant.WebHostEnvironment),
	)
}
