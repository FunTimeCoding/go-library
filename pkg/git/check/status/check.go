package status

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git/check/status/option"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"strings"
)

func Check(o *option.Status) {
	elements := collect(3)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	for _, r := range elements {
		if r.IsClean {
			continue
		}

		fmt.Printf("%s\n", r.Path)

		if !r.IsClean && r.Status != "" {
			fmt.Printf(
				"  Changes:\n%s\n",
				strings.TrimSpace(r.Status),
			)
		}

		fmt.Println()
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
