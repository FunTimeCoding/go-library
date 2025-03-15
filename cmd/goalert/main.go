package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
)

func main() {
	c := alertmanager.NewEnvironment()

	for _, a := range c.Alerts() {
		fmt.Println(a.Format(format.ExtendedColor))
	}
}
