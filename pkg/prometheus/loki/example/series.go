package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/prometheus/loki"
)

func Series() {
	c := loki.NewEnvironment(false)
	fmt.Printf("Series: %s\n", c.Series(`{namespace="bot"}`))
}
