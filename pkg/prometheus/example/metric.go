package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus"
)

func Metric() {
	c := prometheus.NewEnvironment()

	for _, m := range c.AllMetrics() {
		fmt.Printf("Metric: %s\n", m)
	}
}
