package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki"
	"time"
)

func Load() {
	c := loki.NewEnvironment()
	end := time.Now()

	if false {
		start := end.AddDate(0, 0, -7)

		for _, l := range c.Labels(start, end) {
			fmt.Printf("Label: %s\n", l)
			fmt.Printf("  Values: %+v\n", c.LabelValues(start, end, l))
		}
	}

	start := end.Add(-1 * time.Hour)
	fmt.Printf(
		"QueryRange: %+v\n",
		c.QueryRange("{namespace=\"i9n\"} |= ``", start, end),
	)

	if false {
		fmt.Printf(
			"Query: %+v\n",
			c.Query("{namespace=\"i9n\"} |= ``"),
		)
		c.Statistic("")
		c.Series("")
	}
}
