package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
)

const (
	Owner      = "funtimecoding"
	Repository = "go-library"
)

func main() {
	c := github.NewEnvironment()

	for _, w := range c.Workflows(Owner, Repository) {
		fmt.Printf("Workflow: %s\n", w.Format())
	}

	for _, r := range c.Runs(Owner, Repository) {
		fmt.Printf("Run: %s\n", r.Format())

		for _, j := range c.Jobs(Owner, Repository, *r.Raw.ID) {
			fmt.Printf("  Run: %s\n", j.Format())
		}
	}
}
