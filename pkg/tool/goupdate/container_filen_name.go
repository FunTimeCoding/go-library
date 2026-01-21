package goupdate

import (
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/system"
)

func ContainerFileName() string {
	if system.FileExists(project.ContainerFile) {
		return project.ContainerFile
	}

	if system.FileExists(project.DockerFile) {
		return project.DockerFile
	}

	return ""
}
