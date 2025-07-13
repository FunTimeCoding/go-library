package project

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
)

func New(
	path string,
	projectVersion string,
	runtimeVersion string,
) *Project {
	return &Project{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			constant.GoPrefix,
			path,
		),
		Path:           path,
		Version:        projectVersion,
		runtimeVersion: runtimeVersion,
	}
}
