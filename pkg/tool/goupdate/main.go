package goupdate

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/github"
	githubConstant "github.com/funtimecoding/go-library/pkg/github/constant"
	"github.com/funtimecoding/go-library/pkg/github/release"
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goupdate/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(argument.Continue, false, "Continue on error")
	var exclusives []string
	a.StringSliceVariable(
		&exclusives,
		argument.Exclusive,
		nil,
		"One or more matches to exclusively update, comma separated",
	)
	var downgrades []string
	a.StringSliceVariable(
		&downgrades,
		argument.Downgrade,
		nil,
		"One or more downgrades to apply after update, comma separated",
	)
	a.Parse(version, gitHash, buildDate)
	continueOnError := a.GetBoolean(argument.Continue)

	if len(exclusives) > 0 {
		fmt.Printf("Exclusive matches: %s\n", join.Comma(exclusives))
	}

	if len(downgrades) > 0 {
		fmt.Printf("Downgrades: %s\n", join.Comma(downgrades))
	}

	go_mod.UpdateDirectDependencies(exclusives, continueOnError)
	go_mod.DowngradeDependencies(downgrades)
	go_mod.Tidy()
	goVersion := runtime.ExecutableVersion()

	if goVersion == nil {
		system.Exitf(1, "could not determine Go version\n")

		return
	}

	goString := goVersion.String()
	system.Run(
		library.Go,
		library.Mod,
		library.Edit,
		key_value.Equals(library.VersionArgument, goString),
	)

	if file := ContainerFileName(); file != "" {
		d := project.ReplaceGoFromVersion(
			system.ReadFile(system.WorkingDirectory(), file),
			goString,
		)

		if token := os.Getenv(githubConstant.TokenEnvironment); token != "" {
			if v := release.LatestCompatible(
				github.New(
					token,
				).MustReleases(
					githubConstant.DelveNamespace,
					githubConstant.DelveRepository,
				),
				goVersion,
			); v != nil {
				d = project.ReplaceDelveVersion(d, semver.Trim(v.Name))
			}
		}

		system.SaveFile(file, d)
	}

	if system.FileExists(project.GitLabFile) {
		d := project.ReplaceGoImageVersion(
			system.ReadFile(system.WorkingDirectory(), project.GitLabFile),
			goString,
		)
		system.SaveFile(project.GitLabFile, d)
	}
}
