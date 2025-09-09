package clean

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/git/remote"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/lab"
	"github.com/funtimecoding/go-library/pkg/tool/goclean/clean/option"
)

func Lab(
	o *option.Clean,
	origin *remote.Remote,
) {
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
	c := gitlab.New(
		o.GitLabHost,
		environment.Get(constant.TokenEnvironment),
	)
	p := c.ProjectByName(namespace, repository)

	if o.Verbose {
		fmt.Printf(
			"Project: %d %s %s\n",
			p.Identifier,
			p.Namespace,
			p.Name,
		)
	}

	lab.Pipeline(o, c, p)
	lab.Registry(c, p)
	lab.Package(c, p)
	lab.Tag(c, p)
}
