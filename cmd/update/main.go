package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/release"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
	"strings"
)

func main() {
	goVersion := runtime.Version()
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
				d = project.ReplaceDelveVersion(
					d,
					strings.TrimPrefix(v.Name, git.VersionPrefix),
				)
			}
		}

		system.SaveFile(project.DockerFile, d)
	}
}
