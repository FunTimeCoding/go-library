package alert

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule"
	"time"
)

func printRules(
	c *alertmanager.Client,
	firing bool,
	old bool,
) {
	f := option.ExtendedColor.Copy()

	for _, r := range c.Rules().Alert() {
		if old &&
			r.RawAlert != nil &&
			time.Since(
				r.RawAlert.LastEvaluation,
			).Round(time.Second) < 1*time.Minute {
			continue
		}

		if firing && r.State != rule.FiringState {
			continue
		}

		fmt.Println(r.Format(f))
	}
}
