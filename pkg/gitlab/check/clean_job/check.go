package clean_job

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func Check() {
	pflag.String(argument.Namespace, "", "Namespace")
	pflag.String(argument.Project, "", "Project")
	pflag.String(argument.Match, "", "Description match")
	argument.ParseBind()
	g := gitlab.NewEnvironment()
	f := constant.Format.Copy().Tag(tag.Dense)
	m := viper.GetString(argument.Match)

	if m == "" {
		fmt.Printf(
			"--%s must match a runner description\n",
			argument.Match,
		)

		for _, r := range g.Runners(true) {
			fmt.Println(r.Format(f))
		}

		os.Exit(1)
	}

	r := g.RunnerByDescriptionMatch(m)

	if r == nil {
		fmt.Println("No runner match")

		os.Exit(1)
	}

	RunnerWay(g, r, f)

	if false {
		p := g.ProjectByName(
			argument.RequiredStringFlag(argument.Namespace),
			argument.RequiredStringFlag(argument.Project),
		)

		if false {
			PipelineWay(g, p, f)
		}

		if false {
			ProjectWay(g, p, f)
		}
	}
}
