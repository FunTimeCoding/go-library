package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/check/alert/option"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
)

func Print(p *option.Alert) {
	c := internal.Opsgenie()

	if p.Notation {
		printNotation(c)

		return
	}

	f := statusOption.ExtendedColor.Copy()
	alerts := c.Open()

	for _, a := range alerts {
		fmt.Println(a.Format(f))
	}

	if len(alerts) == 0 {
		fmt.Println("No relevant alerts")
	}
}
