package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Export() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	for _, p := range c.ChildPages(c.DefaultSpace(), c.DefaultPage()) {
		fmt.Println(p.Format(f))
		c.Export(
			p,
			fmt.Sprintf("fixture/wiki/example/%s.json", p.Name),
		)
	}
}
