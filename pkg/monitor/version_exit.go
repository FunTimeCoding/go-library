package monitor

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/spf13/viper"
	"os"
)

func VersionExit(
	version string,
	gitHash string,
	buildDate string,
) {
	if !viper.GetBool(argument.Version) {
		return
	}

	b := build.New(version, gitHash, buildDate)
	b.Print()
	os.Exit(0)
}
