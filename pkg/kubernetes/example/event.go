package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kubernetes/client"
	"github.com/funtimecoding/go-library/pkg/kubernetes/constant"
	core "k8s.io/api/core/v1"
)

func Event() {
	c := client.NewEnvironment()
	f := constant.Format

	for _, n := range c.Events(
		core.NamespaceAll,
		0,
		constant.TypeWarning,
	) {
		fmt.Println(n.Format(f))
	}
}
