package skip_configuration

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"slices"
)

func New(
	raw string,
	verbose bool,
) *Configuration {
	result := &Configuration{Raw: raw}

	if raw != "" {
		result.Skips = split.Comma(raw)
	}

	for _, skip := range []string{
		constant.Directory,
		systemConstant.IdeaPath,
	} {
		withSlash := fmt.Sprintf("%s/", skip)

		if !slices.Contains(result.Skips, withSlash) {
			result.Skips = append(result.Skips, withSlash)
		}
	}

	result.Count = len(result.Skips)

	if result.Count > 0 && verbose {
		fmt.Printf(
			"Skips (%d): %s\n",
			result.Count,
			join.Comma(result.Skips),
		)
	}

	return result
}
