package gmail

import (
	"github.com/funtimecoding/go-library/pkg/gmail/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Client {
	return New(
		environment.Default(
			constant.DirectoryEnvironment,
			system.Join(system.Home(), constant.DefaultDirectory),
		),
	)
}
