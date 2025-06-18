package fetch

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"strings"
)

func list() []string {
	result := []string{
		constant.GoAlert,
		constant.GoFile,
		constant.GoGenie,
		constant.GoJira,
		constant.GoKevt,
		constant.GoSentry,
		constant.GoSilence,
	}

	if s := environment.GetDefault(
		constant.PluginEnvironment,
		"",
	); s != "" {
		if strings.HasPrefix(s, separator.Plus) {
			result = append(result, split.Comma(s)...)
		} else {
			result = split.Comma(s)
		}
	}

	return result
}
