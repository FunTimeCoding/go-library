package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func Page() {
	c := confluence.NewEnvironment()
	f := constant.DefaultFormat

	for _, p := range c.Pages() {
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
