package goclaude

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/command_context"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	var host string
	var port int
	c := command_context.New()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			c.Initialize(host, port)
		},
	}
	o.PersistentFlags().StringVar(
		&host,
		"host",
		environment.Fallback(constant.HostEnvironment, web.Localhost),
		"goclauded host",
	)
	o.PersistentFlags().IntVar(
		&port,
		"port",
		environment.FallbackInteger(constant.PortEnvironment, web.ListenPort),
		"goclauded port",
	)
	o.AddCommand(sessionBranch(c))
	o.AddCommand(register(c))
	o.AddCommand(check(c))
	o.AddCommand(wait(c))
	errors.PanicOnError(o.Execute())
}
