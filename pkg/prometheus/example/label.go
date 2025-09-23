package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus"
)

func Label() {
	fmt.Println("Label")

	for _, m := range prometheus.NewEnvironment().AllLabels() {
		fmt.Printf("  %s\n", m)
	}
}
