package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func Page() {
	c := confluence.NewEnvironment()
	s := c.DefaultSpace()
	p := c.DefaultPage()
	f := constant.Dense

	if false {
		a := c.PageBySpaceAndName(s, "Delta")

		if a != nil {
			c.Delete(a.Identifier)
		}

		c.Import(s, p, "fixture/wiki/example/Delta.json")
	}

	if false {
		for _, g := range c.PagesBySpace(c.SpaceByName(s).Identifier) {
			fmt.Println(g.Format(f))

			if false {
				if g.Name != p {
					continue
				}
			}

			if false {
				page.PrintBody(g.Raw.Body)
			}
		}
	}

	if false {
		for _, g := range c.Pages() {
			fmt.Println(g.Format(f))
		}
	}
}
