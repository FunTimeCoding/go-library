package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
)

func Registry() {
	c := gitlab.NewEnvironment()
	p := c.MustProjectByName("shiin", "go-mint")
	repos := c.MustRegistryRepositories(p.Identifier, false)
	fmt.Printf("Repos: %d\n", len(repos))

	for _, r := range repos {
		fmt.Printf("Registry: %d %s\n", r.ID, r.Name)
	}
}
