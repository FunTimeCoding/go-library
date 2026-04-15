package gonetb

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/netb"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gonetb", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	c := netb.NewEnvironment()
	root := &cobra.Command{
		Use:     "gonetb",
		Version: fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate),
	}
	root.AddCommand(listDevicesCommand(c))
	root.AddCommand(getDeviceCommand(c))
	root.AddCommand(listInterfacesCommand(c))
	root.AddCommand(listAddressesCommand(c))
	root.AddCommand(listSitesCommand(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
