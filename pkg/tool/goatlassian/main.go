package goatlassian

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassian/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/client"
	"github.com/spf13/cobra"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Name, version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	c := client.NewEnvironment()
	o := &cobra.Command{
		Use:     constant.Name,
		Version: argument.CobraVersion(version, gitHash, buildDate),
	}
	o.AddCommand(searchIssues(c))
	o.AddCommand(getIssue(c))
	o.AddCommand(listProjects(c))
	o.AddCommand(searchPages(c))
	o.AddCommand(getPage(c))
	o.AddCommand(createPage(c))
	o.AddCommand(updatePage(c))
	o.AddCommand(listSpaces(c))
	o.AddCommand(getPageChildren(c))
	o.AddCommand(getTransitions(c))
	o.AddCommand(transitionIssue(c))
	o.AddCommand(addIssueComment(c))
	o.AddCommand(addPageComment(c))
	errors.PanicOnError(o.Execute())
}
