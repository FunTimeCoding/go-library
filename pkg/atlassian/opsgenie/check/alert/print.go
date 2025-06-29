package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
)

func Print(o *option.Alert) {
	c := internal.Opsgenie()
	alerts := c.Open()

	if o.Notation {
		printNotation(alerts, o)

		return
	}

	f := constant.Format

	for _, a := range alerts {
		fmt.Println(a.Format(f))
	}

	if len(alerts) == 0 {
		fmt.Println("No relevant alerts")
	}
}
