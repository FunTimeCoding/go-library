package raid_parser

import (
	"github.com/funtimecoding/go-library/pkg/raid_parser/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(environment.Required(constant.HostEnvironment))
}
