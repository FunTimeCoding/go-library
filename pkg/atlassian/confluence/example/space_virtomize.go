package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
)

func SpaceVirtomize() {
	c := confluence.NewEnvironment()
	fmt.Printf("%+v\n", c.UserVirtomize())

	for _, s := range c.SpacesVirtomize() {
		fmt.Println(s.Name)
	}
}
