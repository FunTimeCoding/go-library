package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
)

func Node() {
	k := client.NewEnvironment()
	f := constant.Format

	for _, n := range k.Nodes() {
		fmt.Println(n.Format(f))
	}
}
