package skip_configuration

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/split"
)

func New(
	raw string,
	verbose bool,
) *Configuration {
	var result []string

	if raw != "" {
		result = split.Comma(raw)
	}

	if c := len(result); c > 0 && verbose {
		fmt.Printf("Skips (%d): %+v\n", c, result)
	}

	return &Configuration{Raw: raw, Skips: result}
}
