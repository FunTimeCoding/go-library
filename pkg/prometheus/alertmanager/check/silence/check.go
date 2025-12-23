package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert/advanced_option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
	"time"
)

func Check(o *option.Silence) {
	c := internal.Alertmanager()

	if o.Notation {
		printNotation(c, o)

		return
	}

	if o.Set != "" {
		fmt.Printf("Set: %s\n", c.SimpleSilence(o.Set))
	}

	o2 := advanced_option.New()
	o2.All = true
	a, _ := c.Alerts(o2)
	fmt.Printf("Alerts: %d\n", len(a))
	silences := c.Silences(true)

	if !o.All {
		silences = silence.FilterPermanent(silences)
	}

	var relevant int
	f := constant.Format
	t := time.Now()

	for _, e := range silences {
		if !o.All && e.State != constant.ActiveState {
			continue
		}

		relevant++
		fmt.Println(e.Format(f))

		if m := matcher.Matches(e, a, t); len(m) > 0 {
			fmt.Printf("  Matching: %d\n", len(m))
		}
	}

	if !o.All && relevant == 0 {
		fmt.Printf(
			"No relevant %s, %d in total\n",
			item.GoSilence.Plural,
			len(silences),
		)
	}
}
