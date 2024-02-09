package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func main() {
	c := gitlab.New(
		environment.Get(gitlab.Host),
		environment.Get(gitlab.Token),
	)
	// TODO: Detect if on GitLab
	//  GitLab:
	//   TODO: Delete pipelines except latest
	//   TODO: Delete registry images except latest and its version
	//   TODO: Delete tags except latest
	// TODO: Detect if on GitHub:
	//   TODO: Delete all tags except the latest
	// TODO: Read project name from origin remote
	p := c.ProjectByName("foo", "bar")
	fmt.Printf("Project: %d\n", p.Identifier)

	for _, element := range c.Pipelines(p.Identifier) {
		fmt.Printf("Pipeline: %+v\n", element)
	}

	// TODO: Why is this forbidden?
	// TODO: Separate between package and docker repository
	for _, element := range c.RegistryRepositories(p.Identifier) {
		fmt.Printf("Registry repository: %+v\n", element)

		for _, inner := range c.Images(p.Identifier, element.ID) {
			fmt.Printf("Image: %+v\n", inner)
		}
	}

	c.Tags(p.Identifier)
}
