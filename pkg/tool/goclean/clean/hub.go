package clean

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/run"
	hubTag "github.com/funtimecoding/go-library/pkg/github/tag"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Hub(origin *remote.Remote) {
	remoteLocator := git.ParseLocator(origin.Locator)

	if remoteLocator == nil {
		system.Exitf(
			1,
			"could not parse remote locator: %s\n",
			origin.Locator,
		)

		return
	}

	namespace, repository := git.ParseProject(remoteLocator.Path)
	c := github.NewEnvironment()
	tags := c.Tags(namespace, repository)
	latestTag := hubTag.Latest(tags)

	if latestTag == nil {
		system.Exitf(
			1,
			"could not find latest tag for %s/%s\n",
			namespace,
			repository,
		)

		return
	}

	for _, t := range tags {
		if t.Name == latestTag.Name {
			continue
		}

		fmt.Printf("Delete tag: %s\n", *t.Name)
		c.DeleteTag(namespace, repository, *t.Name)
	}

	git.Fetch()

	runs := c.ProjectRuns(namespace, repository)

	if false {
		latestRun := run.Latest(runs)
		fmt.Printf("Latest run: %s\n", latestRun.Name)
	}

	for _, r := range runs {
		if r.Status != run.Completed {
			continue
		}

		fmt.Printf("Delete run: %d\n", r.Identifier)
		c.DeleteRun(namespace, repository, r.Identifier)
	}
}
