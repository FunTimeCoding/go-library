package example

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence"

func Page() {
	c := confluence.NewEnvironment()
	s := c.DefaultSpace()
	p := c.DefaultPage()

	if false {
		if a := c.PageBySpaceAndName(s, "Charlie"); a != nil {
			c.Delete(a.Identifier)
		}

		c.Import(s, p, "fixture/wiki/example/Charlie.json")
	}
}
