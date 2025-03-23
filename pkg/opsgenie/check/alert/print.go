package alert

import (
	"fmt"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/opsgenie"
	"github.com/funtimecoding/go-library/pkg/opsgenie/check/alert/option"
)

func Print(p *option.Alert) {
	c := opsgenie.NewEnvironment()

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
