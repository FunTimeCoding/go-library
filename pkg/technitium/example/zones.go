package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/technitium"
)

func Zones() {
	c := technitium.NewEnvironment()

	for _, z := range c.MustZones() {
		fmt.Printf("%s (%s)\n", z.Name, z.Type)
	}
}
