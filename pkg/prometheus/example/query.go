package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/maps"
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/parse"
	"time"
)

func Query() {
	c := prometheus.NewEnvironment()
	now := time.Now()

	for _, k := range maps.StringKeys(c.QueryIntegers("up", now)) {
		fmt.Printf("Up: %s\n", k)
	}

	countPerScrapeJob := c.QueryIntegers(
		`count by (job)({__name__=~".+"})`,
		now,
	)

	for _, k := range maps.StringKeys(countPerScrapeJob) {
		fmt.Printf(
			"Scrape Job: %s Count: %d\n",
			k,
			countPerScrapeJob[k],
		)
	}

	cardinalityPerMetric := c.QueryIntegers(
		`count by (__name__)({__name__=~".+"})`,
		now,
	)

	for _, k := range maps.StringKeys(cardinalityPerMetric) {
		fmt.Printf(
			"Metric: %s Count: %d\n",
			k,
			cardinalityPerMetric[k],
		)
	}

	load1 := c.QueryFloat(constant.Load1, now)
	load5 := c.QueryFloat(constant.Load5, now)
	load15 := c.QueryFloat(constant.Load15, now)
	fmt.Printf("Load: %.1f %.1f %.1f\n", load1, load5, load15)

	for _, r := range parse.Generic(c.Query(constant.Load1, now)) {
		fmt.Printf("  %s %s %s\n", r.Metric, r.Time, r.Value)
	}
}
