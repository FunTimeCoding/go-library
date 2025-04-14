package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
)

func Print(o *option.Alert) {
	c := internal.Opsgenie()

	if o.Notation {
		printNotation(c, o)

		return
	}

	f := constant.Format
	alerts := c.Open()

	for _, a := range alerts {
		fmt.Println(a.Format(f))
	}

	if len(alerts) == 0 {
		fmt.Println("No relevant alerts")
	}
}
