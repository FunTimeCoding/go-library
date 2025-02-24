package build

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func GuessBinaryPath(
	name string,
	systemArchitecture string,
) string {
	if s := system.Join(
		constant.Temporary,
		name,
		systemArchitecture,
		name,
	); system.FileExists(s) {
		return s
	}

	return ""
}
