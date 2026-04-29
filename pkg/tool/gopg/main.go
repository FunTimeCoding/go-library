package gopg

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/client"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gopg", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	address := environment.Optional("GOPGD_ADDRESS")

	if address == "" {
		address = "http://localhost:8080"
	}

	v, e := client.NewClient(address)
	errors.PanicOnError(e)
	root := &cobra.Command{
		Use:     "gopg",
		Version: fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate),
	}
	root.AddCommand(listInstancesCommand(v))
	root.AddCommand(queryCommand(v))
	root.AddCommand(explainCommand(v))
	root.AddCommand(listSchemasCommand(v))
	root.AddCommand(listTablesCommand(v))
	root.AddCommand(describeTableCommand(v))
	root.AddCommand(listIndexesCommand(v))
	root.AddCommand(tableSizesCommand(v))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
