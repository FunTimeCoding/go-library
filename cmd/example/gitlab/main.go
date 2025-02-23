package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	if true {
		CloneAll()
	}

	if false {
		job.Check()
	}
}

func CloneAll() {
	g := gitlab.NewEnvironment()
	base := system.Join(system.Home(), "gitlab-backup")
	projects := g.Projects()
	count := len(projects)
	var i int

	for _, p := range projects {
		i++
		fmt.Printf(
			"Project (%d/%d): %s\n",
			i,
			count,
			p.PathWithNamespace,
		)
		group := system.Join(base, p.Namespace.Path)
		system.EnsurePathExists(group)
		repository := system.Join(group, p.Path)
		fmt.Printf("  Clone %s to %s\n", p.SSHURLToRepo, repository)

		if !system.DirectoryExists(repository) {
			system.ChangeDirectory(group)
			git.Run("clone", p.SSHURLToRepo)

			if false {
				// Fails with SSH agent on Windows
				git.Clone(p.SSHURLToRepo, repository)
			}
		}
	}
}
