package gohabitica

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gohabitica/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/client"
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
	c := client.NewEnvironment()
	root := &cobra.Command{
		Use:     constant.Name,
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	root.AddCommand(tasks(c))
	root.AddCommand(create(c))
	root.AddCommand(score(c))
	root.AddCommand(tags(c))
	root.AddCommand(statistic(c))
	root.AddCommand(cron(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
