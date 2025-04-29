package issue

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"

func (i *Issue) Option() *option.Issue {
	return i.option
}
