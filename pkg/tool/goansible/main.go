package goansible

import "github.com/funtimecoding/go-library/pkg/monitor"

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	// TODO: Find Ansible API, allow calling it from CI/CD
}
