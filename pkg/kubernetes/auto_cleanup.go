package kubernetes

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func AutoCleanup() bool {
	return environment.Fallback(
		constant.AutoCleanupEnvironment,
		"",
	) != ""
}
