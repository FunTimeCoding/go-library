package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
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

	silences := c.Silences(true)
	var relevant int
	f := constant.Format

	for _, e := range silences {
		if !o.All && e.State != constant.ActiveState {
			continue
		}

		relevant++
		fmt.Println(e.Format(f))
	}

	if !o.All && relevant == 0 {
		fmt.Printf(
			"No relevant %s, %d in total\n",
			Plural,
			len(silences),
		)
	}
}
