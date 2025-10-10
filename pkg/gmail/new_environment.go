package gmail

import (
	"github.com/funtimecoding/go-library/pkg/gmail/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func NewEnvironment() *Client {
	return New(
		environment.Fallback(
			constant.DirectoryEnvironment,
			join.Absolute(system.Home(), constant.DefaultDirectory),
		),
	)
}
