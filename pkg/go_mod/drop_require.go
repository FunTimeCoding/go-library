package go_mod

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func dropRequire(mod string) {
	system.Run(
		constant.Go,
		constant.Mod,
		constant.Edit,
		fmt.Sprintf("-droprequire=%s", mod),
	)
}
