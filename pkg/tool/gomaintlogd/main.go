package gomaintlogd

import (
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/relational/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	monitor.ParseBind(version, gitHash, buildDate)
	o := option.New()
	o.PostgresLocator = environment.Required(postgres.LocatorEnvironment)
	o.Version = version
	Run(o)
}
