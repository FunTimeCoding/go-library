package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func Node() {
	c := client.NewEnvironment()
	f := constant.Format

	for _, n := range c.Nodes() {
		fmt.Println(n.Format(f))
	}
}
