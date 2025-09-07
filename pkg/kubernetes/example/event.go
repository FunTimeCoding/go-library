package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func Event() {
	k := client.NewEnvironment()
	f := constant.Format

	for _, n := range k.EventsSimple(false, true) {
		fmt.Println(n.Format(f))
	}
}
