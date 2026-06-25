package gotechnitium

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/technitium"
	"github.com/funtimecoding/go-library/pkg/tool/gotechnitium/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := technitium.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(listZones(c))
	o.AddCommand(createZone(c))
	o.AddCommand(listRecords(c))
	o.AddCommand(addRecord(c))
	o.AddCommand(deleteRecord(c))
	errors.PanicOnError(o.Execute())
}
