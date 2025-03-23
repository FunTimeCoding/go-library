package silence

import (
	"fmt"
	statusOption "github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/option"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func Print(p *option.Silence) {
	c := alertmanager.NewEnvironment()

	if p.Notation {
		printNotation(c)

		return
	}

	if p.Set != "" {
		fmt.Printf("Set: %s\n", c.SimpleSilence(p.Set))
	}

	silences := c.Silences(true)
	var relevant int
	f := statusOption.Color.Copy().Tag(tag.Link)

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
