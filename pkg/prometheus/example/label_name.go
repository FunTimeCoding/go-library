package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus"
)

func LabelName() {
	c := prometheus.NewEnvironment()

	for _, l := range c.LabelNames([]string{}, constant.StartOfTime) {
		fmt.Printf("Label: %s\n", l)
	}
}
