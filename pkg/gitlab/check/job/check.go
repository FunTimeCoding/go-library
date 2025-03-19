package job

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/spf13/pflag"
)

func Check() {
	pflag.String(argument.Namespace, "", "Namespace")
	pflag.String(argument.Project, "", "Project")
	pflag.String(argument.Match, "", "Description match")
	argument.ParseAndBind()
	g := gitlab.NewEnvironment()

	if true {
		RunnerWay(g, argument.RequiredStringFlag(argument.Match, 1))
	}

	if false {
		p := g.ProjectByName(
			argument.RequiredStringFlag(argument.Namespace, 1),
			argument.RequiredStringFlag(argument.Project, 1),
		)

		if false {
			PipelineWay(g, p)
		}

		if false {
			ProjectWay(g, p)
		}
	}

	// TODO: Retry if failed, but only if not already retried and successful afterwards
}
