package issue

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

func fixtureOption() *option.Issue {
	return option.New(
		"http://localhost",
		"test",
		[]string{},
		[]string{constant.Closed},
		field_map.New([]jira.Field{}),
	)
}
