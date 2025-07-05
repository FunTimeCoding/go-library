package git

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/system"
)

func IsCleanCommand() bool {
	return system.Run(
		constant.Command,
		constant.Status,
		constant.Porcelain,
	) == ""
}
