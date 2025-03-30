package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
)

func Page() {
	c := confluence.NewEnvironment()
	fmt.Println("PagesBasic")
	fmt.Println(c.PagesBasic())
}
