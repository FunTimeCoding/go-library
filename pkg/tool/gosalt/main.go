package gosalt

import "github.com/funtimecoding/go-library/pkg/monitor"

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	// TODO: Use Salt API, allow calling it from CI/CD
}
