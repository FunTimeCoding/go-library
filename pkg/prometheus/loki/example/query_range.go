package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki/basic/constant"
	timeLibrary "github.com/funtimecoding/go-library/pkg/time"
	telemetry "github.com/funtimecoding/go-library/pkg/web/telemetry/constant"
	"strings"
	"time"
)

func QueryRange() {
	c := loki.NewEnvironment(false)
	end := time.Now()
	start := end.Add(-2 * time.Hour)
	r, m := c.QueryRange(
		`{namespace="bot"} | json | msg="request_start"`,
		start,
		end,
	)

	if false {
		fmt.Printf("QueryRange: %s %+v\n", m.Type, m.Statistic)
	}

	for _, v := range r {
		if false {
			if v.Stream != constant.Stdout {
				continue
			}
		}

		route := v.ValueString(telemetry.Route)
		body := v.ValueString(telemetry.Body)

		if strings.HasPrefix(route, "/test-") {
			continue
		}

		fmt.Printf(
			"%s %s %s %s\n",
			v.Time.Format(timeLibrary.DateMinute),
			console.Cyan("%s", route),
			v.Stream,
			body,
		)

		if v.Text != "" {
			fmt.Printf("  Text: >%s<\n", v.Text)
		}
	}
}
