package project

import "github.com/funtimecoding/go-library/pkg/monitor/item/constant"

func New(
	path string,
	projectVersion string,
	runtimeVersion string,
) *Project {
	return &Project{
		MonitorIdentifier: constant.GoVersion.StringIdentifier(path),
		Path:              path,
		Version:           projectVersion,
		runtimeVersion:    runtimeVersion,
	}
}
