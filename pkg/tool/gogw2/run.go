package gogw2

import (
	"github.com/funtimecoding/go-library/pkg/gw2/check/alliance"
	"github.com/funtimecoding/go-library/pkg/tool/gogw2/option"
)

func Run(o *option.Alliance) {
	alliance.Check()

	if false {
		alliance.PrintAccount()
	}
}
