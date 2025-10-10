package build

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func GuessBinaryPath(
	name string,
	systemArchitecture string,
) string {
	if s := join.Relative(
		constant.Temporary,
		name,
		systemArchitecture,
		name,
	); system.FileExists(s) {
		return s
	}

	return ""
}
