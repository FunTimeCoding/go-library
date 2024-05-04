package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/release"
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	var exclusives []string
	pflag.StringSliceVar(
		&exclusives,
		argument.Exclusive,
		nil,
		"One or more matches to exclusively update, comma separated",
	)
	var downgrades []string
	pflag.StringSliceVar(
		&downgrades,
		argument.Downgrade,
		nil,
		"One or more downgrades to apply after update, comma separated",
	)
	argument.ParseAndBind()

	fmt.Printf("Exclusive matches: %+v\n", exclusives)
	fmt.Printf("Downgrades: %+v\n", downgrades)
	go_mod.UpdateDirectDependencies(exclusives)
	go_mod.DowngradeDependencies(downgrades)

	goVersion := runtime.ExecutableVersion()
	goString := goVersion.String()
	system.Run("go", "mod", "edit", fmt.Sprintf("-go=%s", goString))

	if system.FileExists(project.DockerFile) {
		d := project.ReplaceGoFromVersion(
			system.ReadFile(project.DockerFile),
			goString,
		)

		if token := os.Getenv(github.Token); token != "" {
			if v := release.LatestCompatible(
				github.New(
					token,
				).Releases(github.DelveNamespace, github.DelveRepository),
				goVersion,
			); v != nil {
				d = project.ReplaceDelveVersion(d, semver.Trim(v.Name))
			}
		}

		system.SaveFile(project.DockerFile, d)
	}

	if system.FileExists(project.GitLabFile) {
		d := project.ReplaceGoImageVersion(
			system.ReadFile(project.GitLabFile),
			goString,
		)

		system.SaveFile(project.GitLabFile, d)
	}
}
