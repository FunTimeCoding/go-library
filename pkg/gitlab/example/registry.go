package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
)

func Registry() {
	c := gitlab.NewEnvironment()
	repositories := c.MustRegistryRepositories(
		c.MustProjectByName("shiin", "go-mint").Identifier,
		false,
	)
	fmt.Printf("Repositories: %d\n", len(repositories))

	for _, r := range repositories {
		fmt.Printf("Registry: %d %s\n", r.ID, r.Name)
	}
}
