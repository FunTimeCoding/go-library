package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	v := runtime.Version().String()
	system.Run("go", "mod", "edit", fmt.Sprintf("-go=%s", v))

	if system.FileExists(project.DockerFile) {
		d := project.ReplaceGoFromVersion(
			system.ReadFile(project.DockerFile),
			v,
		)
		d = project.ReplaceDelveVersion(d, v)
		system.SaveFile(project.DockerFile, d)
	}
}
