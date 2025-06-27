package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
)

func Search() {
	c := github.NewEnvironment()

	if false {
		for _, r := range c.SearchRepository(
			"user:%s",
			c.User().Name,
		) {
			fmt.Printf("Repository: %s\n", r.Name)
		}
	}

	if true {
		for _, r := range c.ActionRepository() {
			fmt.Printf("Code: %+v\n", *r.Raw.Name)
		}
	}
}
