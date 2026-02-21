package gochk

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/gochk/check"
	"github.com/spf13/pflag"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	pflag.String(
		argument.Port,
		"",
		"Port, multiple values separated by comma",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	check.Check()
}
