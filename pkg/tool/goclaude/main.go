package goclaude

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	o := &cobra.Command{
		Use:     constant.Identity.Usage(),
		Short:   constant.Identity.Description(),
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	c, e := client.NewClientWithResponses(
		locator.New("localhost").Port(8583).Insecure().String(),
	)
	errors.PanicOnError(e)
	o.AddCommand(sessionBranch(c))
	o.AddCommand(register())
	o.AddCommand(check())
	o.AddCommand(wait())
	errors.PanicOnError(o.Execute())
}
