package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/opsgenie/alert/prometheus_detail"
)

func parseDescription(
	d string,
	verbose bool,
) *prometheus_detail.Detail {
	if verbose {
		fmt.Printf("  Description: %s\n", console.Magenta(d))
	}

	result := prometheus_detail.New()

	// TODO: What is generic here?

	return result
}
