package forge

import (
	"github.com/funtimecoding/go-library/pkg/forge/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func AutoCleanup() bool {
	return environment.Exists(constant.AutoCleanupEnvironment)
}
