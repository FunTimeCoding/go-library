package gohab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/client"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("gohab", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	c := client.NewEnvironment()
	root := &cobra.Command{
		Use:     "gohab",
		Version: fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate),
	}
	root.AddCommand(tasksCommand(c))
	root.AddCommand(createCommand(c))
	root.AddCommand(scoreCommand(c))
	root.AddCommand(tagsCommand(c))
	root.AddCommand(statsCommand(c))
	root.AddCommand(cronCommand(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
