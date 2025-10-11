package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
	webConstant "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func fixtureOption() *option.Issue {
	return option.New(
		locator.New(webConstant.Localhost).Insecure().String(),
		"test",
		[]string{},
		[]string{constant.Closed},
		field_map.New([]jira.Field{}),
	)
}
