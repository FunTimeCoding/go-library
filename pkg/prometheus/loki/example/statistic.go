package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki"
)

func Statistic() {
	c := loki.NewEnvironment(true)
	// TODO: Somehow just returns zeroes
	fmt.Printf("Statistic: %s", c.Statistic(`{namespace!=""}`))
}
