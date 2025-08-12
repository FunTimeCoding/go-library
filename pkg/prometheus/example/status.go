package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus"
)

func Status() {
	c := prometheus.NewEnvironment()

	fmt.Printf("Status: %+v\n", c.Status())
}
