package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/salt"
)

func Minions() {
	c := salt.NewEnvironment()
	result, e := c.Minions()

	if e != nil {
		fmt.Printf("error: %v\n", e)

		return
	}

	for _, m := range result {
		fmt.Printf("%s (%s %s)\n", m.ID, m.OS, m.OSRelease)
	}
}
