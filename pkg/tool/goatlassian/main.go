package goatlassian

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassian/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/client"
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
	root.AddCommand(searchIssues(c))
	root.AddCommand(getIssue(c))
	root.AddCommand(listProjects(c))
	root.AddCommand(searchPages(c))
	root.AddCommand(getPage(c))
	root.AddCommand(createPage(c))
	root.AddCommand(updatePage(c))
	root.AddCommand(listSpaces(c))
	root.AddCommand(getPageChildren(c))
	root.AddCommand(getTransitions(c))
	root.AddCommand(transitionIssue(c))
	root.AddCommand(addIssueComment(c))
	root.AddCommand(addPageComment(c))

	if f := root.Execute(); f != nil {
		errors.Printf("%v", f)
		os.Exit(1)
	}
}
