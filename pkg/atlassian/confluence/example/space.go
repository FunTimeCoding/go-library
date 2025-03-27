package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
)

func Space() {
	c := confluence.NewEnvironment()
	u := c.User()
	fmt.Println(u)

	for _, s := range c.Spaces() {
		fmt.Println(s.Name)
	}
}
