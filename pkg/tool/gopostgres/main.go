package gopostgres

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgres/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/generated/client"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	v, e := client.NewClient(
		locator.New(
			environment.Fallback(
				constant.HostEnvironment,
				web.Localhost,
			),
		).Port(web.ListenPort).Insecure().String(),
	)
	errors.PanicOnError(e)
	root := &cobra.Command{
		Use:     constant.Name,
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	root.AddCommand(listInstances(v))
	root.AddCommand(query(v))
	root.AddCommand(explain(v))
	root.AddCommand(listSchemas(v))
	root.AddCommand(listTables(v))
	root.AddCommand(describeTable(v))
	root.AddCommand(listIndexes(v))
	root.AddCommand(tableSizes(v))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
