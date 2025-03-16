package silence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/check/silence/parameter"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"
)

func Print(p *parameter.Silence) {
	if p.Notation {
		printNotation()

		return
	}

	c := alertmanager.NewEnvironment()

	if p.Set != "" {
		fmt.Printf("Set: %s\n", c.SimpleSilence(p.Set))
	}

	silences := c.Silences(true)
	var relevant int
	f := format.Color.Copy().Tag(tag.Link)

	for _, a := range silences {
		if !p.All && a.State != constant.ActiveState {
			continue
		}

		relevant++
		fmt.Println(a.Format(f))
	}

	if !p.All && relevant == 0 {
		fmt.Printf("No relevant silences, %d in total\n", len(silences))
	}
}
