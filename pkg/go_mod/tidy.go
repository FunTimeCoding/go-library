package go_mod

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Tidy() {
	system.Run(constant.Go, constant.Mod, constant.Tidy)
}
