package issues

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/basic_client/issue"

func (i *Issues) Key(k string) *issue.Issue {
	for _, s := range i.list {
		if s.Key == k {
			return s
		}
	}

	return nil
}
