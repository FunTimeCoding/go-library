package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Required(constant.UserKeyEnvironment),
		environment.Required(constant.TeamKeyEnvironment),
		environment.Required(constant.WebHostEnvironment),
	)
}
