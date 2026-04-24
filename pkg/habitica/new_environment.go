package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
