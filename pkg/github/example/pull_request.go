package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
)

func PullRequest() {
	g := github.NewEnvironment()
	u := g.User()
	repositories := g.Repositories(u.Name)

	for _, r := range repositories {
		fmt.Printf("Repository: %s\n", *r.Name)

		for _, p := range g.PullRequests(u.Name, *r.Name) {
			fmt.Printf("  PR: %s\n", *p.Title)
			fmt.Printf("  %s\n", *p.HTMLURL)
		}
	}

	if true {
		for _, r := range repositories {
			fmt.Printf("Repository: %s\n", *r.Name)

			for _, i := range g.ProjectIssues(u.Name, *r.Name) {
				fmt.Printf("  Issue: %s\n", *i.Title)
				fmt.Printf("  %s\n", *i.HTMLURL)
			}
		}
	}
}
