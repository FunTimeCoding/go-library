package goatl

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atl"
	"github.com/funtimecoding/go-library/pkg/errors"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
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
		r := reporter.New("goatl", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	c := atl.NewEnvironment()
	root := &cobra.Command{
		Use:     "goatl",
		Version: fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate),
	}
	root.AddCommand(searchIssuesCommand(c))
	root.AddCommand(getIssueCommand(c))
	root.AddCommand(listProjectsCommand(c))
	root.AddCommand(searchPagesCommand(c))
	root.AddCommand(getPageCommand(c))
	root.AddCommand(createPageCommand(c))
	root.AddCommand(updatePageCommand(c))
	root.AddCommand(listSpacesCommand(c))
	root.AddCommand(getPageChildrenCommand(c))
	root.AddCommand(getTransitionsCommand(c))
	root.AddCommand(transitionIssueCommand(c))
	root.AddCommand(addIssueCommentCommand(c))
	root.AddCommand(addPageCommentCommand(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
