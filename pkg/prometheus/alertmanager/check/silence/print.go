package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func Print(p *option.Silence) {
	c := internal.Alertmanager()

	if p.Notation {
		printNotation(c)

		return
	}

	if p.Set != "" {
		fmt.Printf("Set: %s\n", c.SimpleSilence(p.Set))
	}

	silences := c.Silences(true)
	var relevant int
	f := constant.Format

	for _, a := range silences {
		if !p.All && a.State != constant.ActiveState {
			continue
		}

		relevant++
		fmt.Println(a.Format(f))
	}

	if !p.All && relevant == 0 {
		fmt.Printf(
			"No relevant silences, %d in total\n",
			len(silences),
		)
	}
}
