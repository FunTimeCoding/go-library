package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
)

func Check(o *option.Alert) {
	elements := collect()

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
