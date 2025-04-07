package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus"
)

func Label() {
	c := prometheus.NewEnvironment()

	for _, l := range c.LabelNames([]string{}, constant.StartOfTime) {
		fmt.Printf("Label: %s\n", l)
	}
}
