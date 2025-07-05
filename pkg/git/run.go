package git

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Run(s ...string) {
	system.Run(append([]string{constant.Command}, s...)...)
}
