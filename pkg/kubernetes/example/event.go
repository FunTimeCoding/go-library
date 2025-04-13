package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func Event() {
	c := client.NewEnvironment()
	f := constant.Format

	for _, n := range c.EventsSimple(false, true) {
		fmt.Println(n.Format(f))
	}
}
