package goraid

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/funtimecoding/go-library/pkg/tool/goraid/constant"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := raid.NewEnvironment()
	o := &cobra.Command{
		Use:     "goraid",
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(logs(c))
	o.AddCommand(reports(c))
	errors.PanicOnError(o.Execute())
}
