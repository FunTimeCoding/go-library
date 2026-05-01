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

	for _, s := range c.MustSpaces() {
		fmt.Println(s.Format(f))
	}

	fmt.Println("Page")

	for _, a := range c.MustPagesBySpaceName(c.DefaultSpace()) {
		fmt.Println(a.Format(f))
	}
}
