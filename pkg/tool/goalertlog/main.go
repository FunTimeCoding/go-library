package goalertlog

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlog/option"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	pflag.String(
		argument.Path,
		"tmp/goalertlog.db",
		"Database path",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.DatabasePath = viper.GetString(argument.Path)
	o.Version = version
	Run(o)
}
