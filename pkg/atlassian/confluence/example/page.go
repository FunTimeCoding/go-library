package example

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func Page() {
	c := confluence.NewEnvironment()
	s := c.DefaultSpace()
	p := c.DefaultPage()

	if false {
		if a := c.MustPageBySpaceAndName(s, strings.Charlie); a != nil {
			c.MustDelete(a.Identifier)
		}

		c.MustImport(s, p, "fixture/wiki/example/", "Charlie.json")
	}
}
