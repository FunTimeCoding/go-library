package example

import (
	"fmt"

	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Overview() {
	c := confluence.NewEnvironment(confluence.WithVerbose(true))
	f := constant.Format
	fmt.Println("Space")

	for _, s := range c.Spaces() {
		fmt.Println(s.Format(f))
	}

	fmt.Println("Page")

	for _, a := range c.PagesBySpaceName(c.DefaultSpace()) {
		fmt.Println(a.Format(f))
	}
}
