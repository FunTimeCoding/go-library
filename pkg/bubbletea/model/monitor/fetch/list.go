package fetch

import (
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"strings"
)

func List() []string {
	var result []string

	for _, c := range item.Collectors {
		result = append(result, c.Name)
	}

	if s := environment.Default(
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
