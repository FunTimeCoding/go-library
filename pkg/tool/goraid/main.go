package goraid

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/spf13/cobra"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New("goraid", version)
	r.Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := raid.NewEnvironment()
	root := &cobra.Command{
		Use:     "goraid",
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	root.AddCommand(logs(c))
	root.AddCommand(reports(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
