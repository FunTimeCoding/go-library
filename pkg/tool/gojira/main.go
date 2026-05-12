package gojira

import (
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/tool/gojira/constant"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Boolean(
		argument.Copyable,
		false,
		"Disable OSC8 links and add a copyable link instead",
	)
	a.Boolean(argument.Notation, false, "JSON output")
	a.Boolean(argument.All, false, "Include filtered in output")
	a.Parse(version, gitHash, buildDate)
	o := option.New()
	o.Notation = a.GetBoolean(argument.Notation)
	o.Copyable = a.GetBoolean(argument.Copyable)
	o.All = a.GetBoolean(argument.All)
	issue.Check(o)
}
