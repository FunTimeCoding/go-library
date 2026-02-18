package silence

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
)

func collect(c *alertmanager.Client) []*silence.Silence {
	return c.Silences(true)
}
