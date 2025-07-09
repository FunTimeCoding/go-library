package file

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func collect() []string {
	result := argument.Positionals()

	if s := environment.GetDefault(
		constant.FileEnvironment,
		"",
	); s != "" {
		result = append(result, split.Comma(s)...)
	}

	return result
}
