package file

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func collect(paths []string) []string {
	result := paths

	if s := environment.Optional(constant.FileEnvironment); s != "" {
		result = append(result, split.Comma(s)...)
	}

	return result
}
