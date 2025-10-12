package action

import (
	"github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func IsActionRun() bool {
	return environment.Fallback(constant.RunEnvironment, "") != ""
}
