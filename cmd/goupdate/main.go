package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/release"
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/semver"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func main() {
	go_mod.UpdateDirectDependencies()
	go_mod.Tidy()
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
