package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Page() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	if true {
		p := c.PageBySpaceAndName(constant.OperationsSpace, "Delta")

		if p != nil {
			c.Delete(p.Identifier)
		}
	}

	if false {
		c.Import(
			constant.OperationsSpace,
			constant.ExamplePage,
			"tmp/wiki/Delta.json",
		)
	}

	if false {
		system.EnsurePathExists("tmp/wiki")

		for _, p := range c.ChildPages(
			constant.OperationsSpace,
			constant.ExamplePage,
		) {
			fmt.Println(p.Format(f))
			c.SaveAsFile(p, fmt.Sprintf("tmp/wiki/%s.json", p.Name))
		}
	}

	if false {
		a := c.PageBySpaceAndName(
			constant.OperationsSpace,
			constant.ExamplePage,
		)
		fmt.Println(a.Format(f))
	}

	if false {
		s := c.SpaceByName(constant.OperationsSpace)

		for _, p := range c.PagesBySpace(s.Identifier) {
			fmt.Println(p.Format(f))

			if false {
				if p.Name != constant.ExamplePage {
					continue
				}
			}

			if false {
				page.PrintBody(p.Raw.Body)
			}
		}
	}

	if false {
		for _, p := range c.Pages() {
			fmt.Println(p.Format(f))
		}
	}
}
