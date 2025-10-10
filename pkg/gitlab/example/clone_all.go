package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func CloneAll() {
	g := gitlab.NewEnvironment()
	base := join.Absolute(system.Home(), "gitlab-backup")
	projects := g.Projects()
	count := len(projects)
	var i int

	for _, p := range projects {
		i++
		fmt.Printf(
			"Project (%d/%d): %s\n",
			i,
			count,
			p.Raw.PathWithNamespace,
		)
		group := join.Absolute(base, p.Raw.Namespace.Path)
		system.EnsurePathExists(group)
		repository := join.Absolute(group, p.Raw.Path)
		fmt.Printf(
			"  Clone %s to %s\n",
			p.Raw.SSHURLToRepo,
			repository,
		)

		if !system.DirectoryExists(repository) {
			system.ChangeDirectory(group)
			git.Run(constant.Clone, p.Raw.SSHURLToRepo)

			if false {
				// Fails with SSH agent on Windows
				git.Clone(p.Raw.SSHURLToRepo, repository)
			}
		}
	}
}
