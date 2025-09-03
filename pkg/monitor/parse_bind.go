package monitor

import "github.com/funtimecoding/go-library/pkg/argument"

func ParseBind(
	version string,
	gitHash string,
	buildDate string,
) {
	VersionArgument()
	argument.ParseBind()
	VersionExit(version, gitHash, buildDate)
}
