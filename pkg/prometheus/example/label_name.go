package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus"
)

func LabelName() {
	fmt.Println("Label names")

	for _, l := range prometheus.NewEnvironment().MustLabelNames(
		[]string{},
		constant.StartOfTime,
	).Values {
		fmt.Printf("  %s\n", l)
	}
}
