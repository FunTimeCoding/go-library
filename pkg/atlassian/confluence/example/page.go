package example

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence"

func Page() {
	c := confluence.NewEnvironment()
	s := c.DefaultSpace()
	p := c.DefaultPage()

	if false {
		a := c.PageBySpaceAndName(s, "Delta")

		if a != nil {
			c.Delete(a.Identifier)
		}

		c.Import(s, p, "fixture/wiki/example/Delta.json")
	}
}
