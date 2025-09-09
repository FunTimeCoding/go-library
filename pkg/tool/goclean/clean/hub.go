package clean

import (
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/hub"
)

func Hub(r *remote.Remote) {
	locator := git.ParseLocator(r.Locator)

	if locator == nil {
		system.Exitf(
			1,
			"could not parse remote locator: %s\n",
			r.Locator,
		)

		return
	}

	namespace, repository := git.ParseProject(locator.Path)
	c := github.NewEnvironment()
	hub.Tag(c, namespace, repository)
	hub.Run(c, namespace, repository)
}
