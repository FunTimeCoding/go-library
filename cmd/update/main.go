package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/system"
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
		// TODO: Find latest version by go minor version (for compatibility)
		//  Example: If go is 1.21.7, find latest 1.21.* Delve
		d = project.ReplaceDelveVersion(d, goString)
		system.SaveFile(project.DockerFile, d)
	}
}
