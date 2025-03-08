package skip_configuration

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/split"
)

func New(
	raw string,
	verbose bool,
) *Configuration {
	result := &Configuration{Raw: raw}

	if raw != "" {
		result.Skips = split.Comma(raw)
		result.Count = len(result.Skips)
	}

	if result.Count > 0 && verbose {
		fmt.Printf(
			"Skips (%d): %s\n",
			result.Count,
			join.Comma(result.Skips),
		)
	}

	return result
}
