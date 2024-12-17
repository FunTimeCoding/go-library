package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/check/job"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
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

	for _, element := range projects {
		i++
		fmt.Printf(
			"Project (%d/%d): %s\n",
			i,
			count,
			element.PathWithNamespace,
		)
		group := system.Join(base, element.Namespace.Path)
		system.EnsurePathExists(group)
		repository := system.Join(group, element.Path)
		fmt.Printf(
			"  Clone %s to %s\n",
			element.SSHURLToRepo,
			repository,
		)

		if !system.DirectoryExists(repository) {
			errors.PanicOnError(os.Chdir(group))
			git.Run("clone", element.SSHURLToRepo)

			if false {
				// Fails with SSH agent on Windows
				git.Clone(element.SSHURLToRepo, repository)
			}
		}
	}
}
