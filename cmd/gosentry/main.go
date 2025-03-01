package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func main() {
	p := monitor.BindFlag()
	c := sentry.NewEnvironment()
	issues := c.AllIssues()

	if p.Notation {
		r := report.New()

		for _, i := range issues {
			r.AddItem(
				i.MonitorIdentifier,
				constant.ErrorType,
				i.Title,
				i.Link,
			)
		}

		r.Print()

		return
	}

	for _, i := range issues {
		fmt.Println(i.Format(format.Color))
	}
}
