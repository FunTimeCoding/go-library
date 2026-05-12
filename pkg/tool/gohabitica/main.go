package gohabitica

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gohabitica/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(tasks(c))
	o.AddCommand(create(c))
	o.AddCommand(score(c))
	o.AddCommand(tags(c))
	o.AddCommand(statistic(c))
	o.AddCommand(cron(c))
	errors.PanicOnError(o.Execute())
}
